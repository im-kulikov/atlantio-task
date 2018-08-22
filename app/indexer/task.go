package indexer

import (
	"context"

	"github.com/im-kulikov/atlantio-task/app/eth"
	"github.com/im-kulikov/atlantio-task/app/store"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type task struct {
	cli    *eth.Client
	cfg    *Config
	log    *zap.SugaredLogger
	store  *store.Store
	chain  chan interface{}
	blocks []*store.Block
	addr   map[string]struct{}
	ebn    int64
	sbn    int64
	tbn    int64

	// global error:
	err error
}

var (
	errNothingToSync    = errors.New("nothing to think")
	errEthNodeOutOfSync = errors.New("eth out of sync")
)

// create new task for scan blocks
func newTask(i *Indexer) *task {
	t := &task{
		cli:    i.cli,
		cfg:    i.cfg,
		chain:  i.chain,
		log:    i.log,
		store:  i.store,
		blocks: make([]*store.Block, 0),
		addr:   make(map[string]struct{}),
		err:    nil,
	}

	for _, addr := range i.cfg.Addresses {
		t.addr[addr] = struct{}{}
	}

	if i == nil {
		t.err = errors.New("empty indexer")
	}

	return t
}

// canScan - check if need to scan
func (t *task) canScan(ctx context.Context) *task {
	if t.err != nil {
		return t
	}

	// check that context not canceled
	select {
	case <-ctx.Done():
		t.err = ctx.Err()
	default: // skip
	}

	var err error

	// get last block in store:
	t.sbn = t.store.LastBlock()
	t.tbn = t.sbn

	if t.ebn, err = t.cli.GetLastBlock(ctx); err != nil {
		t.err = errors.WithMessage(err, "eth_getBlockByNumber")
		return t
	}

	if t.ebn > 0 && t.sbn == t.ebn {
		t.err = errNothingToSync
		return t
	}

	if t.ebn < t.sbn {
		t.err = errEthNodeOutOfSync
		return t
	}

	return t
}

// sendData to indexer-chain
func (t *task) sendData(v interface{}) error {
	select {
	case t.chain <- v:
		// data sent successfully
		return nil
	default:
		t.err = errors.New("can't store to chain")
	}
	return t.err
}

// scanBlocks from ethereum
func (t *task) scanBlocks(ctx context.Context) *task {
	if t.err != nil {
		return t
	}

	for {
		// if context canceled - we must stop our work
		select {
		case <-ctx.Done():
			t.err = ctx.Err()
			return t
		default: // we can continue..
		}

		if t.chain == nil {
			t.err = errors.New("indexer-chain is closed")
			return t
		}

		// try to fetch block from ethereum
		block, err := t.cli.GetBlockByNumber(ctx, t.tbn)
		if err != nil {
			t.err = err
			return t
		}

		if block.Number > 0 {
			// store block
			t.blocks = append(t.blocks, block)
		} else if block.Number < 0 {
			return t
		}

		t.tbn++
	}

	return t
}

// storeData - sync blocks and txs
func (t *task) storeData() *task {
	for _, block := range t.blocks {
		if err := t.sendData(block); err != nil {
			return t
		}

		t.tbn = block.Number

		for _, tx := range block.Transactions {
			// if from addr is our - sync it
			if _, ok := t.addr[tx.From]; ok {
				// --------------------------
				if err := t.sendData(tx); err != nil {
					return t
				}
			}

			// if to addr is our - sync it
			if _, ok := t.addr[tx.To]; ok {
				// --------------------------
				if err := t.sendData(tx); err != nil {
					return t
				}
			}
		}
	}

	return t
}

func (t *task) logResult() {
	switch t.err {
	case nil: // ignore
	case context.Canceled: // ignore
	case errNothingToSync: // ignore
	case eth.ErrBlockNotFound: // ignore
	default:
		t.log.Errorw("job failed",
			"new_blocks", t.tbn-t.sbn,
			"store_block", t.sbn,
			"eth_block", t.ebn,
			"last_block", t.tbn,
			"cause", t.err)
		return
	}

	// store block number
	if err := t.sendData(t.tbn); err != nil {
		t.log.Errorw("can't save last block number",
			"error", err)
	}

	t.log.Infow("scan job result",
		"new_blocks", t.tbn-t.sbn,
		"store_block", t.sbn,
		"eth_block", t.ebn,
		"last_block", t.tbn,
		"cause", t.err)
}
