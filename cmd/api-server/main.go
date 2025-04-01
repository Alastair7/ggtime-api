package main

import (
	"github.com/Alastair7/ggtime-api/internal/common"
	"github.com/Alastair7/ggtime-api/internal/server"
)

func main() {
	common.LoadEnvironmentVariables()
	serverConfig := server.NewServerConfiguration()
	server := server.NewApiServer(serverConfig)

	server.RunServer()
}
