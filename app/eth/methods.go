package eth

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/im-kulikov/atlantio-task/app/store"
	"github.com/pkg/errors"
)

type (
	transactionResponse struct {
		Hash        string `json:"hash"`
		From        string `json:"from"`
		To          string `json:"to"`
		BlockHash   string `json:"blockHash"`
		BlockNumber ethInt `json:"blockNumber"`
	}

	blockResponse struct {
		Number       ethInt                 `json:"number"`
		Hash         string                 `json:"hash"`
		Transactions []*transactionResponse `json:"transactions,omitempty"`
		Timestamp    ethTimestamp           `json:"timestamp"`
	}

	signResponse struct {
		Raw string `json:"raw"`
		// we are not interested in other fields
		// Tx *struct {
		// 	BlockHash        interface{} `json:"blockHash"`
		// 	BlockNumber      interface{} `json:"blockNumber"`
		// 	ChainID          string      `json:"chainId"`
		// 	Condition        interface{} `json:"condition"`
		// 	Creates          interface{} `json:"creates"`
		// 	From             string      `json:"from"`
		// 	Gas              string      `json:"gas"`
		// 	GasPrice         string      `json:"gasPrice"`
		// 	Hash             string      `json:"hash"`
		// 	Input            string      `json:"input"`
		// 	Nonce            string      `json:"nonce"`
		// 	PublicKey        string      `json:"publicKey"`
		// 	R                string      `json:"r"`
		// 	Raw              string      `json:"raw"`
		// 	S                string      `json:"s"`
		// 	StandardV        string      `json:"standardV"`
		// 	To               string      `json:"to"`
		// 	TransactionIndex interface{} `json:"transactionIndex"`
		// 	V                string      `json:"v"`
		// 	Value            string      `json:"value"`
		// } `json:"tx"`
	}

	SendParams struct {
		From  string
		To    string
		Value float64
	}
)

var one = big.NewFloat(1e18)

func (b *blockResponse) convert() *store.Block {
	bl := &store.Block{
		Number:    b.Number.Value,
		BlockTime: b.Timestamp.Value,
		CreatedAt: b.Timestamp.Value,
		UpdatedAt: b.Timestamp.Value,
	}

	for _, tx := range b.Transactions {
		bl.Transactions = append(bl.Transactions, &store.Transaction{
			Hash:        tx.Hash,
			From:        tx.From,
			To:          tx.To,
			CreatedAt:   b.Timestamp.Value,
			UpdatedAt:   b.Timestamp.Value,
			Seen:        false,
			BlockNumber: b.Number.Value,
		})
	}

	return bl
}

// GetBlockByNumber - fetch block and transactions from ethereum blockchain
func (c *Client) GetBlockByNumber(ctx context.Context, num int64) (*store.Block, error) {
	var result = new(blockResponse)

	result.Number.Value = -1

	numStr := ethNumFromInt(num)

	err := c.call(callParams{
		ctx:     ctx,
		method:  "eth_getBlockByNumber",
		params:  []interface{}{numStr, true}, // https://wiki.parity.io/JSONRPC-eth-module#eth_getblockbynumber
		result:  result,
		timeout: c.cfg.BlockTimeout,
	})

	if err != nil {
		return nil, err
	}

	if result.Number.Value == -1 {
		return nil, ErrBlockNotFound
	}

	// convert blockResponse to store.Block with transactions
	return result.convert(), nil
}

func (c *Client) GetLastBlock(ctx context.Context) (int64, error) {
	var result = new(ethInt)

	err := c.call(callParams{
		ctx:     ctx,
		method:  "eth_blockNumber",
		params:  []interface{}{}, // https://wiki.parity.io/JSONRPC-eth-module#eth_blocknumber
		result:  result,
		timeout: c.cfg.BlockTimeout,
	})

	if err != nil {
		return 0, err
	}

	// convert eth response to int64
	return result.Value, nil
}

func (c *Client) SendTransaction(ctx context.Context, params SendParams) (string, error) {
	var (
		signResult = new(signResponse)
		sendResult string
		amount     = new(big.Int)
	)

	// Convert amount to wei
	fAmount := big.NewFloat(params.Value)
	fAmount.Mul(fAmount, one).Int(amount)

	signParams := map[string]string{
		"from":     params.From,
		"to":       params.To,
		"gas":      "0x5208",     // 21000
		"gasPrice": "0x3b9aca00", // 1 gWei
		"value":    hexutil.EncodeBig(amount),
		"data":     "0x",
	}

	if err := c.call(callParams{
		ctx:     ctx,
		method:  "eth_signTransaction",
		params:  []interface{}{signParams},
		result:  signResult,
		timeout: c.cfg.SendTransactionTimeout,
	}); err != nil {
		return "", err
	}

	if len(signResult.Raw) == 0 {
		return "", errors.New("signer data is empty")
	}

	if err := c.call(callParams{
		ctx:     ctx,
		method:  "eth_sendRawTransaction",
		params:  []interface{}{signResult.Raw},
		result:  &sendResult,
		timeout: c.cfg.SendTransactionTimeout,
	}); err != nil {
		return "", err
	}

	return sendResult, nil
}
