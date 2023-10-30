package main

import (
	"doctor/internal/config"
	"doctor/internal/db"
	"doctor/internal/service"
	"log"
)

func init() {
	config, err := config.LoadConfig("./")

	if err != nil {
		panic(err)
	}

	err = db.ConnectDB(config)

	if err != nil {
		panic(err)
	}

	err = db.ConnectRedis(config)

	if err != nil {
		panic(err)
	}
}

func main() {
	log.Println("Starting Grpc Server")

	_, err := service.StartGrpcService(&config.Konfig)

	if err != nil {
		log.Println(err)
	}

	log.Println("Grpc Server Running")
}
