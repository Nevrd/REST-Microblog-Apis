package main

import (
	"API/internal/http"
	"context"
)

func main() {
	dataBaseContext, dataBaseContextCancel := context.WithCancel(context.Background())
	defer dataBaseContextCancel()
	server, err := http.NewServer(dataBaseContext)
	if err != nil {
		panic(err)
	}
	if err = server.StartServer(); err != nil {
		panic(err)
	}
}
