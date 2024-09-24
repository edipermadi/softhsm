package main

import (
	"github.com/edipermadi/softhsm/pkg/token/users"
	"log"
	"net"

	"github.com/edipermadi/softhsm/pkg/slip"
	"github.com/edipermadi/softhsm/pkg/token/servers"
	"github.com/edipermadi/softhsm/pkg/token/services"
)

func main() {
	l, err := net.Listen("tcp4", ":12345")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		userService, err := users.NewService()
		if err != nil {
			panic(err)
		}

		reader := slip.NewReader(conn)
		writer := slip.NewWriter(conn)
		buffer := make([]byte, 1024)
		svc := services.NewService(userService)
		svr := servers.NewServer(svc)
		for {
			length, err := reader.Read(buffer)
			if err != nil {
				panic(err)
			}
			payload, err := svr.Process(buffer[0:length])
			if err != nil {
				panic(err)
			}
			if _, err := writer.Write(payload); err != nil {
				panic(err)
			}
		}
	}
}
