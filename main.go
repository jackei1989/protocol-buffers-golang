package main

import (
	"flag"
	"fmt"
	"github.com/jackei1989/protocol-buffers-v2-golang/communication"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

func main() {
	op := flag.String("op", "s", "s for Server and c for Client.")
	flag.Parse()

	switch strings.ToLower(*op) {
	case "s":
		fmt.Println("server is runnnig...")
		RunServer()
	case "c":
		fmt.Println("Data has been sent.")
		RunClient()
	}
}

func RunServer() {
	l, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		defer c.Close()
		go func(cn net.Conn) {
			defer c.Close()
			data, err := ioutil.ReadAll(c)
			if err != nil {
				log.Fatal(err)
			}
			p := &communication.Book{}
			proto.Unmarshal(data, p)

			fmt.Println(p)
		}(c)

	}
}

func RunClient() {
	p := communication.Book{
		Id:        proto.Int64(1),
		Title:     proto.String("Tamaris"),
		Author:    proto.String("George Sand"),
		Published: proto.String("1862"),
	}
	data, err := proto.Marshal(&p)
	if err != nil {
		log.Fatalln(err)
	}

	sendData(data)
}

func sendData(data []byte) {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	write, err := conn.Write(data)
	if err != nil {
		return
	}
	_ = write
}
