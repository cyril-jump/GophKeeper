package main

import (
	"net/http"

	"github.com/cyril-jump/gophkeeper/internal/server/app/server"
)

func main() {

	srv := server.Init()

	if err := srv.Start(":8080"); err != nil && err != http.ErrServerClosed {
		srv.Logger.Fatal(err)
	}
}
