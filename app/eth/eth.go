package eth

import (
	"bytes"
	"net/http"
	"sync"
	"time"

	"github.com/im-kulikov/helium/module"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type (
	Config struct {
		Address                string
		Debug                  bool
		BlockTimeout           time.Duration
		SendTransactionTimeout time.Duration
	}

	Client struct {
		cfg *Config
		cli *http.Client
		log *zap.SugaredLogger
		buf *sync.Pool
		req *http.Request
	}
)

var (
	Module = module.Module{
		{Constructor: NewDefaultConfig},
		{Constructor: NewClient},
	}

	ErrBlockNotFound = errors.New("block not found")
)

const (
	defaultTimeout = time.Second * 5

	charsetUTF8 = "charset=UTF-8"

	MIMEApplicationJSON            = "application/json"
	MIMEApplicationJSONCharsetUTF8 = MIMEApplicationJSON + "; " + charsetUTF8
)

func NewDefaultConfig(v *viper.Viper) (*Config, error) {
	if !v.IsSet("eth") {
		return nil, errors.New("eth: empty config")
	}

	v.SetDefault("eth.block_timeout", defaultTimeout)
	v.SetDefault("eth.send_timeout", defaultTimeout)

	return &Config{
		Address:                v.GetString("eth.address"),
		Debug:                  v.GetBool("eth.debug"),
		BlockTimeout:           v.GetDuration("eth.block_timeout"),
		SendTransactionTimeout: v.GetDuration("eth.send_timeout"),
	}, nil
}

func NewClient(cfg *Config, log *zap.SugaredLogger) (*Client, error) {
	cli := &http.Client{
		// HTTP_PROXY
		// HTTPS_PROXY
		// etc..
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}

	req, err := http.NewRequest(http.MethodPost, cfg.Address, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "eth")
	}
	req.Header.Set("Content-Type", MIMEApplicationJSONCharsetUTF8)
	req.Header.Set("Accept", MIMEApplicationJSON)

	return &Client{
		cfg: cfg,
		cli: cli,
		log: log,
		req: req,
		buf: &sync.Pool{New: func() interface{} {
			return new(bytes.Buffer)
		}},
	}, nil
}
