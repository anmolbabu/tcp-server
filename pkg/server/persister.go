package server

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/anmolbabu/tcp-server/pkg/utils"
)

type Persister struct {
	*Server
	JSONStore *utils.JsonStore
	FileName  string
}

func NewPersister(s Server, jsonStore *utils.JsonStore, fileName string) *Persister {
	p := &Persister{&s, jsonStore, fileName}

	return p
}

func (p *Persister) FetchBufferedData(isRemoveAfterRead bool) (bData utils.SizeStringSlice, err error) {
	_, err = p.JSONStore.Read(&bData, isRemoveAfterRead)
	if err != nil {
		return bData, fmt.Errorf("failed to read data from buffer. Error :%+v\n", err)
	}
	return
}

func (p *Persister) HandleMessage(msg []byte) (err error) {
	// buffer it
	_, err = p.JSONStore.Write(string(msg))
	if err != nil {
		return fmt.Errorf("failed to write message %s to buffer. Error: %+v", err)
	}

	fmt.Println("Successfully stored the data to buffer")
	data, _ := p.FetchBufferedData(false)
	fmt.Println(data.Data, len(data.Data))
	return
}

func (p *Persister) HandleRequest(listener net.Listener) (err error) {
	conn, err := listener.Accept()
	if err != nil {
		return fmt.Errorf("failed accepting requests. Error: %+v", err)
	}

	// Validate message as a valid JSON
	decoder := json.NewDecoder(conn)
	jsonMsg := make(map[string]interface{})
	err = decoder.Decode(&jsonMsg)
	if err != nil {
		return fmt.Errorf("failed to validate message as json. error: %+v", err)
	}

	jsonString, err := json.Marshal(jsonMsg)
	if err != nil {
		return fmt.Errorf("failed to marshal the passed message %+v. Error %+v", jsonMsg, err)
	}

	conn.Write([]byte(fmt.Sprintf("Message received %+v", string(jsonString))))

	err = p.HandleMessage(jsonString)
	if err != nil {
		conn.Close()
		return err
	}

	conn.Close()
	return
}

func (p *Persister) SaveToFile() error {
	f, err := os.OpenFile(p.FileName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("failed to open file %s.Error %+v", p.FileName, err)
	}

	defer f.Close()

	data := utils.NewSizeStringSlice()

	_, err = p.JSONStore.Read(data, true)
	if err != nil {
		return fmt.Errorf("failed to read data in file. Error %+v", err)
	}

	if len(data.Data) > 0 {
		if _, err = f.WriteString(strings.Join(data.Data, "\n")); err != nil {
			if err != nil {
				// Restore data to in memory buffer
				p.JSONStore.WriteMany(data.Data)
				fmt.Printf("failed to write data to file. Error %+v\n", err)
				// Return error
				return fmt.Errorf("failed to write data to file. Error %+v", err)
			}
		}
		f.WriteString("\n")
	}
	return nil
}
