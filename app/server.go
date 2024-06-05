package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type Request struct {
	Method  string
	Path    string
	Headers map[string]string
	Body    string
}

func handleget(conn net.Conn, request Request) {
	defer conn.Close()
	err := error(nil)
	if request.Path == "/" {
		_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
		if err != nil {
			fmt.Println("Error writing response:", err.Error())
			return
		}
	} else if request.Path[0:6] == "/echo/" {
		echo := request.Path[6:]
		fileEncoding := " "
		if len(lines) > 2 {
			if lines[2] != "" {
				if strings.Split(lines[2], ": ")[1] == "gzip" {
					fileEncoding = "gzip"
				} else {
					if len(strings.Split(lines[2], ": ")[1]) > 4 {
						list := strings.Split(strings.Split(lines[2], ": ")[1], ", ")
						for i := 0; i < len(list); i++ {
							if list[i] == "gzip" {
								fileEncoding = "gzip"
								break
							}
						}
					}
				}
			}
		}
		if fileEncoding != " " {
			buffer := new(bytes.Buffer)
			writer := gzip.NewWriter(buffer)
			_, err = writer.Write([]byte(echo))
			if err != nil {
				fmt.Println("Error while writing ", err.Error())
			}
			writer.Close()
			_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: " + fmt.Sprintf("%d", len(buffer.Bytes())) + "\r\nContent-Encoding: " + fileEncoding + "\r\n\r\n" + buffer.String()))
		} else {
			_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: " + fmt.Sprintf("%d", len(echo)) + "\r\n\r\n" + echo))
		}
		if err != nil {
			fmt.Println("Error writing response:", err.Error())
			return
		}
	} else if request.Path[0:7] == "/files/" {
		filepath := path[7:]
		dir := os.Args[2]
		data, err := os.ReadFile(dir + filepath)
		if err != nil {
			_, err = conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
			if err != nil {
				fmt.Println("Error writing response:", err.Error())
				return
			}
		} else {
			_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length: " + strconv.Itoa(len(data)) + "\r\n\r\n" + string(data)))
			if err != nil {
				fmt.Println("Error writing response:", err.Error())
				return
			}
		}
	} else if path == "/user-agent" {
		userAgent := strings.Split(lines[2], ": ")[1]
		_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: " + fmt.Sprintf("%d", len(userAgent)) + "\r\n\r\n" + userAgent))
		if err != nil {
			fmt.Println("Error writing response:", err.Error())
			return
		}
	} else {
		_, err = conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		if err != nil {
			fmt.Println("Error writing response:", err.Error())
			return
		}
	}
}

func handlepost(conn net.Conn, path string) {
	defer conn.Close()
	if path[0:7] == "/files/" {
		filepath := path[7:]
		dir := os.Args[2]
		data := strings.Trim(strings.Split(string(buf), "\r\n\r\n")[1], "\x00")
		err := os.WriteFile(dir+filepath, []byte(data), 0644)
		if err != nil {
			fmt.Println("Error writing into file", err.Error())
			return
		} else {
			_, err = conn.Write([]byte("HTTP/1.1 201 Created\r\n\r\n"))
			if err != nil {
				fmt.Println("Error writing response:", err.Error())
				return
			}
		}
	}
}

func handleconnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}
	fmt.Println("Received data: ", string(buf))
	lines := strings.Split(string(buf), "\r\n")
	request := Request{
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

	fmt.Println("Request: ", request.Method)
	fmt.Println("Path: ", request.Path)
	if request.Method == "GET" {
		handleget(conn)
	} else {
		handlepost(conn)
	}
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
		handleconnection(conn)

	}
}
