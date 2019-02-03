package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:9090")
	// read in input from stdin
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Text to send: ")
	// text, _ := reader.ReadString('\n')
	p := map[string]interface{}{
		"Hi": "World",
	}

	s, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("error marshalling %+v.error %+v\n", p, err)
	}
	// send to socket
	fmt.Fprintf(conn, string(s))
	// listen for reply
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message from server: " + message)

	time.Sleep(5 * time.Second)
	conn1, _ := net.Dial("tcp", "127.0.0.1:8080")
	searchStr := "World"
	fmt.Fprintf(conn1, string(searchStr)+"\n")
	message1, _ := bufio.NewReader(conn1).ReadString('\n')
	fmt.Print("\nMessage from server: " + message1 + "\n")

}
