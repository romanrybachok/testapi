package main

import (
	"github.com/romanrybachok/testapi"
	"github.com/romanrybachok/testapi/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(api.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}
