package main

import (
	config "eulabs/products/etc"
	"eulabs/products/internal/app"
)

func main() {
	println("Main...")
	cfg, err := config.Load("config", "../../")
	if err != nil {
		panic(err)
	}
	a := app.NewApp(cfg)
	err = a.Run()
	if err != nil {
		panic(err)
	}
}
