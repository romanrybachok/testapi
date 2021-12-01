package main

import (
	"github.com/romanrybachok/testapi"
	"github.com/romanrybachok/testapi/pkg/handler"
	"github.com/romanrybachok/testapi/pkg/repository"
	"github.com/romanrybachok/testapi/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(api.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}
