package server

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/anmolbabu/tcp-server/pkg/utils"
)

type Fetcher struct {
	Server
	JSONStore *utils.JsonStore
	FileName  string
}

func NewFetcher(s Server, jsonStore *utils.JsonStore, fileName string) *Fetcher {
	f := &Fetcher{s, jsonStore, fileName}
	return f
}

func (fetcher Fetcher) HandleMessage() (msgs []string, err error) {
	var fileMsgs []string
	bufferedMsgs := utils.NewSizeStringSlice()
	var wg sync.WaitGroup

	wg.Add(2)

	// Read from file and append them to
	go func() {
		defer wg.Done()
		fileMsgs, err = utils.ReadFromFile(fetcher.FileName)
		// Ignore file not found since that is a valid case on start and until the async file storage interval occurs
		if err != nil && !os.IsNotExist(err) {
			err = fmt.Errorf("failed fetching messages from file storage. error: %+v", err)
			return
		}
	}()

	go func() {
		defer wg.Done()
		_, err = fetcher.JSONStore.Read(bufferedMsgs, false)
		if err != nil {
			err = fmt.Errorf("failed fetching messages from buffered store. error: %+v", err)
			return
		}
	}()

	wg.Wait()

	msgs = append(msgs, fileMsgs...)
	msgs = append(msgs, bufferedMsgs.Data...)

	return
}

func (fetcher Fetcher) HandleRequest(listener net.Listener) (err error) {
	conn, err := listener.Accept()
	defer conn.Close()
	if err != nil {
		err = fmt.Errorf("failed accepting requests. Error: %+v", err)
		conn.Write([]byte(err.Error()))
		return
	}

	msg := make([]byte, 1024)
	conn.Read(msg)

	msgs, err := fetcher.HandleMessage()
	if err != nil {
		conn.Write([]byte(err.Error()))
		return err
	}
	fmt.Println(msgs)
	conn.Write([]byte(fmt.Sprintf("Messages received as of now: %+v", strings.Join(msgs, " "))))

	return
}

func (fetcher Fetcher) FetchBufferedData(isRemoveAfterRead bool) (bData utils.SizeStringSlice, err error) {
	return
}
