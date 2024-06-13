package main

import (
	"codecrafters-http-server-go/app/models"
	"codecrafters-http-server-go/app/router"
	"fmt"
	"net"
	"os"
	"strings"
)

func handleconnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}
	lines := strings.Split(string(buf), "\r\n")
	request := models.Request{
		Method: strings.Trim(strings.Split(lines[0], " ")[0], " "),
		Path:   strings.Split(lines[0], " ")[1],
		Headers: func() map[string]string {
			headers := make(map[string]string)
			for _, header := range lines[1:] {
				parts := strings.Split(header, ":")
				if len(parts) == 2 {
					key := strings.TrimSpace(parts[0])
					value := strings.TrimSpace(parts[1])
					headers[key] = value
				}
			}
			return headers
		}(),
		Body: strings.Trim(strings.Split(string(buf), "\r\n\r\n")[1], "\x00"),
	}
	router.Route(conn, request)
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4222")
		os.Exit(1)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		go handleconnection(conn)
	}
}
