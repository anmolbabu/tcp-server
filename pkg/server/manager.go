package server

import (
	"fmt"
	"net"
	"time"

	"github.com/anmolbabu/tcp-server/pkg/config"
	"github.com/anmolbabu/tcp-server/pkg/utils"
)

type ServerInterface interface {
	HandleRequest(net.Listener) error
	FetchBufferedData(bool) (utils.SizeStringSlice, error)
}

type Server struct {
	ConnectionType string
	HostName       string
	Port           int
}

func NewServer(cServer config.Server) Server {
	return Server{
		ConnectionType: cServer.Type,
		HostName:       cServer.Hostname,
		Port:           cServer.Port,
	}
}

func (server Server) Create() (listener net.Listener, err error) {
	listener, err = net.Listen(server.ConnectionType, fmt.Sprintf("%s:%d", server.HostName, server.Port))
	if err != nil {
		return listener, fmt.Errorf("failed to start fetcher server")
	}
	fmt.Printf("Listening on %s:%d\n", server.HostName, server.Port)
	return
}

func (server Server) FetchBufferedData(isRemoveAfterRead bool, js *utils.JsonStore) (bData utils.SizeStringSlice, err error) {
	return
}
func (p Server) HandleRequest(listener net.Listener) (err error) {
	return
}

func (server Server) Cleanup(listener net.Listener) (err error) {
	listener.Close()
	return
}

func Init(conf config.Config, jsonStore *utils.JsonStore) (err error) {
	errCh := make(chan error)

	fServer := NewFetcher(NewServer(conf.Fetcher), jsonStore, conf.DataFile)
	fetcherListener, err := fServer.Create()
	if err != nil {
		return fmt.Errorf("failed to initialise the fetcher server. Error %+v", err)
	}
	defer fServer.Cleanup(fetcherListener)
	go func() {
		for {
			err = fServer.HandleRequest(fetcherListener)
			if err != nil {
				errCh <- fmt.Errorf("failed to handle request on fetcher server %+v. Error : %+v", fServer, err)
			}
		}
	}()

	pServer := NewPersister(NewServer(conf.Listener), jsonStore, conf.DataFile)
	persisterListener, err := pServer.Create()
	if err != nil {
		return fmt.Errorf("failed to initialise the persister server. Error %+v", err)
	}
	defer pServer.Cleanup(persisterListener)
	go func() {
		for {
			err = pServer.HandleRequest(persisterListener)
			if err != nil {
				errCh <- fmt.Errorf("failed to handle request on persister server %+v. Error : %+v", pServer, err)
			}
		}
	}()

	for {
		select {
		case err := <-errCh:
			fmt.Println(err)
		case <-time.After(time.Duration(conf.FileDumpInterval) * time.Second):
			err = pServer.SaveToFile()
			if err != nil {
				return err
			}
		}
	}

	return
}
