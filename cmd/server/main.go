package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/gommon/log"

	"github.com/cyril-jump/gophkeeper/internal/server/app/server"
	"github.com/cyril-jump/gophkeeper/internal/server/pkg/auth/strict"
	"github.com/cyril-jump/gophkeeper/internal/server/pkg/config"
	"github.com/cyril-jump/gophkeeper/internal/server/pkg/provider/postgres"
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

	db := postgres.New(cfg.DatabaseDSN)

	auth := strict.New(ctx)

	srv := server.Init(ctx, db, auth)

	go func() {

		<-signalChan

		log.Info("Shutting down...")

		cancel()
		if err = srv.Shutdown(ctx); err != nil && err != ctx.Err() {
			srv.Logger.Fatal(err)
		}

		if err = db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err = srv.Start(cfg.ServerAddress); err != nil && err != http.ErrServerClosed {
		srv.Logger.Fatal(err)
	}
}
