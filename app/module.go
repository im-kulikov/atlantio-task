package app

import (
	"github.com/im-kulikov/atlantio-task/app/api"
	"github.com/im-kulikov/atlantio-task/app/eth"
	"github.com/im-kulikov/atlantio-task/app/indexer"
	"github.com/im-kulikov/atlantio-task/app/store"
	"github.com/im-kulikov/helium/grace"
	"github.com/im-kulikov/helium/logger"
	"github.com/im-kulikov/helium/module"
	"github.com/im-kulikov/helium/orm"
	"github.com/im-kulikov/helium/redis"
	"github.com/im-kulikov/helium/settings"
	"github.com/im-kulikov/helium/web"
	"github.com/im-kulikov/helium/workers"
)

var Module = module.Module{
	{Constructor: NewApplication},
	{Constructor: NewJobs},
	{Constructor: store.NewStore},
}.
	Append(settings.Module).   // Viper settings
	Append(grace.Module).      // graceful context
	Append(logger.Module).     // Zap Logger
	Append(orm.Module).        // PostgreSQL connection
	Append(workers.Module).    // Workers module
	Append(web.ServersModule). // Web-servers module
	Append(redis.Module).      // Redis connection
	Append(eth.Module).        // Eth client
	Append(indexer.Module).    // Indexer worker
	Append(api.Module)         // API module
