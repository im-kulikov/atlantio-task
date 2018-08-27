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
}.Append(
	settings.Module,   // Viper settings
	grace.Module,      // graceful context
	logger.Module,     // Zap Logger
	orm.Module,        // PostgreSQL connection
	workers.Module,    // Workers module
	web.ServersModule, // Web-servers module
	redis.Module,      // Redis connection
	eth.Module,        // Eth client
	indexer.Module,    // Indexer worker
	api.Module,        // API module
)
