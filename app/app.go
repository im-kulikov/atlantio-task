package app

import (
	"context"

	"github.com/chapsuk/mserv"
	"github.com/chapsuk/worker"
	"github.com/im-kulikov/helium"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

type (
	app struct {
		log     *zap.Logger
		workers *worker.Group
		servers mserv.Server
	}

	Params struct {
		dig.In

		Loggger *zap.Logger
		Workers *worker.Group
		Servers mserv.Server
	}
)

func NewApplication(params Params) helium.App {
	return app{
		log:     params.Loggger,
		workers: params.Workers,
		servers: params.Servers,
	}
}

func (a app) Run(ctx context.Context) error {
	a.log.Info("running workers")
	a.workers.Run()
	a.log.Info("running servers")
	a.servers.Start()

	a.log.Info("app successfully runned")
	<-ctx.Done()

	a.log.Info("stopping http servers...")
	a.servers.Stop()
	a.log.Info("stopping workers...")
	a.workers.Stop()

	a.log.Info("gracefully stopped")
	return nil
}
