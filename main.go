package main

import (
	"cleanArch/wire"
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	app, err := wire.InitializeApplication(ctx)
	if err != nil {
		log.Fatal("Failed initialize app", err.Error())
	}
	app.HttpServer.ListenAndServe()
}
