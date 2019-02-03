package main

import (
	"fmt"

	"github.com/anmolbabu/tcp-server/pkg/server"
	"github.com/anmolbabu/tcp-server/pkg/utils"
)

func main() {
	jsonStore := utils.NewJsonStore()
	err := server.Init(
		server.Server{
			ConnectionType: server.ConnectionType,
			HostName:       "",
			Port:           8080,
		},
		server.Server{
			ConnectionType: server.ConnectionType,
			HostName:       "",
			Port:           9090,
		},
		jsonStore,
		"data.json",
	)
	if err != nil {
		fmt.Println(err)
	}
	return
}
