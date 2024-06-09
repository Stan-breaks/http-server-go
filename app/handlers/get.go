package handlers

import (
	"bytes"
	"codecrafters-http-server-go/app/models"
	"codecrafters-http-server-go/app/utils"
	"compress/gzip"
	"fmt"
	"net"
	"os"
)

func Get(conn net.Conn, request models.Request) {
	err := error(nil)
	if request.Path == "/" {
		err = utils.WriteResponse(conn, 200, nil, "")
		if err != nil {
			fmt.Println("Error writing response:", err.Error())
			return
		}
	} else if request.Path[0:6] == "/echo/" {
		echo := request.Path[6:]
		fileEncoding, ok := request.Headers["Accept-Encoding"]
		if ok {
			buffer := new(bytes.Buffer)
			writer := gzip.NewWriter(buffer)
			_, err = writer.Write([]byte(echo))
			if err != nil {
				fmt.Println("Error while writing ", err.Error())
			}
			writer.Close()
			headers := map[string]string{
				"Content-Encoding": fileEncoding,
			}
			err = utils.WriteResponse(conn, 200, headers, buffer.String())
		} else {
			err = utils.WriteResponse(conn, 200, nil, echo)
		}
		if err != nil {
			fmt.Println("Error writing response:", err.Error())
			return
		}
	} else if request.Path[0:7] == "/files/" {
		filepath := request.Path[7:]
		dir := os.Args[2]
		data, err := os.ReadFile(dir + "/" + filepath)
		if err != nil {
			err = utils.WriteResponse(conn, 404, nil, "")
			if err != nil {
				fmt.Println("Error writing response:", err.Error())
				return
			}
		} else {
			headers := map[string]string{
				"Content-Type": "application/octet-stream",
			}
			err = utils.WriteResponse(conn, 200, headers, string(data))
			if err != nil {
				fmt.Println("Error writing response:", err.Error())
				return
			}
		}
	} else if request.Path == "/user-agent" {
		userAgent := request.Headers["User-Agent"]
		headers := map[string]string{
			"Content-Type": "text/plain",
		}
		err = utils.WriteResponse(conn, 200, headers, userAgent)
		if err != nil {
			fmt.Println("Error writing response:", err.Error())
			return
		}
	} else {
		err = utils.WriteResponse(conn, 404, nil, "")
		if err != nil {
			fmt.Println("Error writing response:", err.Error())
			return
		}
	}
}
