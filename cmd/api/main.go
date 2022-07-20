package main

import (
	"context"
	"log"

	"go-micro.dev/v4/api"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := api.NewApi()

	if err := srv.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
