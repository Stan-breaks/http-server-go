package utils

import (
	"fmt"
	"net"
)

func WriteResponse(conn net.Conn, statusCode int, headers map[string]string, body string) error {
	stringHeaders := ""
	if headers != nil {
		for key, value := range headers {
			stringHeaders += (key + ": " + value + "\r\n")
		}
	}
	statusText, ok := HTTPStatusCodes[statusCode]
	if !ok {
		statusText = fmt.Sprintf("%d Status", statusCode)
	}
	err := error(nil)
	if stringHeaders != "" {
		_, err = conn.Write([]byte("HTTP/1.1 " + statusText + "\r\n\r\n" + body))
	} else {
		_, err = conn.Write([]byte("HTTP/1.1 " + statusText + "\r\n\r\n" + body))
	}

	if err != nil {
		return err
	} else {
		return nil
	}
}
