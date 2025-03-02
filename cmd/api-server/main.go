package main

import (
	"github.com/Alastair7/ggtime-api/internal/server"
)

func main() {
	server := &server.ApiServer{
		Address: ":8080",
	}

	server.RunServer()
}
