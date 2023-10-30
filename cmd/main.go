package main

import (
	"doctor/internal/config"
	"doctor/internal/db"
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

func main() {}
