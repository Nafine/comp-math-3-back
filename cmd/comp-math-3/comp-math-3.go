package main

import (
	"comp-math-3/internal/config"
	"comp-math-3/internal/web"
)

func main() {
	cfg, err := config.Get()

	if err != nil {
		panic(err)
	}

	server := web.New(cfg)

	_ = server.Start()
}
