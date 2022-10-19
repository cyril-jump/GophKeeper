package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/gommon/log"

	"github.com/cyril-jump/gophkeeper/internal/client/app/client"
	"github.com/cyril-jump/gophkeeper/internal/client/pkg/config"
	"github.com/cyril-jump/gophkeeper/internal/client/pkg/types"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	var cfg config.Config
	err := cfg.Parse()
	if err != nil {
		log.Fatal(err)
	}

	cln := client.Init(ctx, cfg)

	err = cln.Run(types.HandlerType(cfg.Request))
	if err != nil {
		log.Fatal(err)
	}

	<-signalChan

	log.Info("Shutting down...")

	cancel()

}
