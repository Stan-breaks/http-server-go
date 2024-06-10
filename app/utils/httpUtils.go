package utils

import (
	"codecrafters-http-server-go/app/models"
	"fmt"
	"net"
)

func WriteResponse(conn net.Conn, statusCode int, headers map[string]string, body string, request models.Request) error {
	stringHeaders := ""

	for key, value := range headers {
		stringHeaders += (key + ": " + value + "\r\n")
	}

	statusText, ok := HTTPStatusCodes[statusCode]
	if !ok {
		statusText = fmt.Sprintf("%d Status", statusCode)
	}
	err := error(nil)

	if body != "" {
		stringHeaders += ("Content-Length: " + fmt.Sprintf("%d", len(body)) + "\r\n")
	}
	if stringHeaders != "" {
		stringHeaders = ("\r\n" + stringHeaders + "\r\n")
		_, err = conn.Write([]byte("HTTP/1.1 " + statusText + stringHeaders + body))
	} else {
		_, err = conn.Write([]byte("HTTP/1.1 " + statusText + "\r\n\r\n"))
	}

	if err != nil {
		return err
	} else {
		fmt.Println(request.Method + "  " + fmt.Sprintf("%d",statusCode) + "  " + request.Path)
		return nil
	}
}
