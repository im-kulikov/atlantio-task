package indexer

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/im-kulikov/atlantio-task/app/eth"
	"github.com/im-kulikov/atlantio-task/app/store"
	"github.com/im-kulikov/helium/module"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

type (
	// Indexer - is a worker, that find transactions in block-chain
	Indexer struct {
		cfg   *Config
		chain chan interface{}
		cli   *eth.Client
		log   *zap.SugaredLogger
		store *store.Store
	}

	// Params to create new worker
	Params struct {
		dig.In

		Context context.Context // graceful.Context
		Client  *eth.Client
		Config  *Config
		Logger  *zap.SugaredLogger
		Store   *store.Store
		Viper   *viper.Viper
	}

	// Config is worker settings
	Config struct {
		Addresses []string
		ChainSize int
		Timeout   time.Duration
	}
)

// Module indexer worker
var Module = module.Module{
	{Constructor: NewConfig},
	{Constructor: NewIndexer},
}

func NewConfig(v *viper.Viper) (*Config, error) {
	var addresses []string

	if !v.IsSet("indexer") {
		return nil, errors.New("indexer: config is empty")
	}

	tmp := v.GetStringSlice("indexer.addresses")

	for _, item := range tmp {
		if !common.IsHexAddress(item) {
			return nil, errors.Errorf("indexer: bad address %q", item)
		}

		addresses = append(addresses, item)
	}

	if !v.IsSet("indexer.addresses") || len(addresses) < 2 {
		return nil, errors.New("indexer: addresses is empty")
	}

	v.SetDefault("indexer.chain_size", 100)
	v.SetDefault("indexer.sync_timeout", time.Second*5)

	return &Config{
		Addresses: addresses,
		ChainSize: v.GetInt("indexer.chain_size"),
		Timeout:   v.GetDuration("indexer.sync_timeout"),
	}, nil
}

// NewIndexer worker
func NewIndexer(params Params) *Indexer {
	i := &Indexer{
		cfg:   params.Config,
		cli:   params.Client,
		log:   params.Logger,
		store: params.Store,
		chain: make(chan interface{}, params.Config.ChainSize),
	}

	// run background job to sync store:
	go i.syncStore(params.Context)

	return i
}

// receive messages from channel and try to sync with store
func (i *Indexer) syncStore(ctx context.Context) {
	td := time.NewTicker(i.cfg.Timeout)
	defer td.Stop()

	for {
		select {
		case <-ctx.Done():
			i.log.Info("indexer: stop sync, context canceled")
			i.chain = nil
			return
		case <-td.C:
			continue // wait for data..
		case v := <-i.chain:
			i.syncData(v)
		}
	}
}

// sync data to store
func (i *Indexer) syncData(v interface{}) {
	switch item := v.(type) {
	case *store.Block:
		if err := i.store.SaveBlock(item); err != nil {
			i.log.Errorw("can't store block", "error", err, "model", item)
		}
	case *store.Transaction:
		if err := i.store.SaveTransaction(item); err != nil {
			i.log.Errorw("can't store transaction", "error", err, "model", item)
		}
	case int64: // block number
		i.store.SetBlock(item)
		i.log.Infow("store block number", "block", item)
	default:
		i.log.Debugw("unknown type", "model", item)
	}
}

// Job to sync blocks and transactions:
func (i *Indexer) Job(ctx context.Context) {
	newTask(i). // create task
			canScan(ctx).    // check internal data
			scanBlocks(ctx). // scan blocks
			storeData().     // store blocks and txs
			logResult()      // log task results
}
