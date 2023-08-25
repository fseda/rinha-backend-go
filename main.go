package main

import (
	"github.com/fseda/rinha-backend-go/app"
)

func main() {
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
