package main

import (
	"vws_api/internal/config"
	"vws_api/internal/routes"
	"vws_api/types"
)

func main() {
	cfg := config.MustLoad()

	VWS := types.VWS{
		Env:      cfg.Env,
		BotToken: cfg.BotToken,
	}

	routes.InitRoutes(&VWS)
}
