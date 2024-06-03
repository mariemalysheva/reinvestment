package main

import (
	"context"
	"fmt"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/container"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/router"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cont := container.New()

	err := cont.RunCronJobs(ctx)
	if err != nil {
		log.Fatalln("failed to init cronjob service", "err", err.Error())
	}

	api, err := cont.GetAPI(ctx)
	if err != nil {
		log.Fatalln("failed to get api", "err", err.Error())
	}

	r := router.NewAPI(api)
	c := cors.New(cors.Options{})

	port := 3000
	log.Printf("Starting server at %d", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), c.Handler(r)); err != nil {
		log.Fatalln("failed to start server", "err", err.Error())
	}
}
