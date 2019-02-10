package main

import (
	"fmt"
	"path/filepath"

	"github.com/anmolbabu/tcp-server/pkg/config"
	"github.com/anmolbabu/tcp-server/pkg/server"
	"github.com/anmolbabu/tcp-server/pkg/utils"
)

func main() {
	jsonStore := utils.NewJsonStore()
	usrHomeDir := utils.GetUserHome()

	confFilePath := filepath.Join(usrHomeDir, ".tcp-server", "tcp-server.conf")
	conf := config.LoadConfiguration(confFilePath)

	err := server.Init(
		conf,
		jsonStore,
	)
	if err != nil {
		fmt.Println(err)
	}
	return
}
