package api

import (
	"context"
	"net/http"
	"time"

	"github.com/im-kulikov/atlantio-task/app/eth"
	"github.com/im-kulikov/atlantio-task/app/indexer"
	"github.com/im-kulikov/atlantio-task/app/store"
	"github.com/im-kulikov/helium/module"
	"github.com/im-kulikov/helium/web"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"go.uber.org/dig"
)

type (
	// Params for API
	Params struct {
		dig.In

		Engine    *echo.Echo
		Client    *eth.Client
		Config    *indexer.Config
		Indexer   *indexer.Indexer
		Store     *store.Store
		Validator web.Validator
		Viper     *viper.Viper
	}

	// sendRequest to process send transaction
	sendRequest struct {
		From   string  `json:"from" validate:"required,address" message:"bad address"`
		To     string  `json:"to" validate:"required,address" message:"bad address"`
		Amount float64 `json:"amount" validate:"required" message:"must be greater than zero"`
	}

	// lastResponse to answer request
	lastResponse struct {
		Date    time.Time `json:"date"`
		Address string    `json:"address"`
		Block   int64     `json:"block,omitempty"`
	}
)

// Module API
var Module = module.Module{
	{Constructor: web.NewEngine},
	{Constructor: web.NewBinder},
	{Constructor: web.NewLogger},
	{Constructor: web.NewValidator},
	{Constructor: NewAPI},
}

// NewAPI creates http.Handler
func NewAPI(params Params) (http.Handler, error) {
	// try to connect validators
	if err := connectValidators(params.Validator); err != nil {
		return nil, err
	}

	var addrHash = make(map[string]struct{}, len(params.Config.Addresses))

	for _, item := range params.Config.Addresses {
		addrHash[item] = struct{}{}
	}

	params.Engine.GET("/api/last", getLast(params.Store, params.Indexer))
	params.Engine.POST("/api/send", sendTransaction(params.Client, addrHash))

	return params.Engine, nil
}

func getLast(db *store.Store, idx *indexer.Indexer) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		txs, err := db.FetchTransactions()
		if err != nil {
			return err // don't wary, this is demo app
		}

		// prepare answer
		result := make([]lastResponse, 0, len(txs))
		toUpdate := make([]*store.Transaction, 0, len(txs))

		for _, tx := range txs {
			result = append(result, lastResponse{
				Date:    tx.CreatedAt,
				Address: tx.To,
				// Block:   tx.BlockNumber, // for debug
			})

			if !tx.Seen {
				tx.Seen = true
				toUpdate = append(toUpdate, tx)
			}
		}

		db.UpdateSeen(toUpdate)

		return ctx.JSON(http.StatusOK, result)
	}
}

// sendTransaction request
func sendTransaction(cli *eth.Client, addresses map[string]struct{}) echo.HandlerFunc {
	bg := context.Background()
	return func(ctx echo.Context) error {
		var (
			req  sendRequest
			err  error
			hash string
		)

		// parse request and validate data
		if err = ctx.Bind(&req); err != nil {
			return err
		}

		// Check that we has this address
		if _, ok := addresses[req.From]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, "from: unknown address")
		}

		// Check that we has this address
		if _, ok := addresses[req.To]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, "to: unknown address")
		}

		// try to send transaction
		if hash, err = cli.SendTransaction(bg, eth.SendParams{
			From:  req.From,
			To:    req.To,
			Value: req.Amount,
		}); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, map[string]string{
			"hash": hash,
		})
	}
}
