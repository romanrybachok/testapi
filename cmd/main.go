package main

import (
	api "github.com/romanrybachok/testapi"
	"log"
)

func main() {
	srv := new(api.Server)

	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}
