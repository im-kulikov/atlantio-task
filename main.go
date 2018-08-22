package main

import (
	"github.com/im-kulikov/atlantio-task/app"
	"github.com/im-kulikov/helium"
	"github.com/im-kulikov/helium/settings"
)

var (
	BuildTime    = "now"
	BuildVersion = "dev"

	cfg = &settings.App{
		File:         "config.yml",
		Name:         "atlant",
		BuildTime:    BuildTime,
		BuildVersion: BuildVersion,
	}
)

func main() {
	h, err := helium.New(cfg, app.Module)
	if err != nil {
		panic(err)
	}

	if err := h.Run(); err != nil {
		panic(err)
	}
}
