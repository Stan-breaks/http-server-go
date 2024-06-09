package handlers

import (
	"codecrafters-http-server-go/app/models"
	"codecrafters-http-server-go/app/utils"
	"bytes"
	"fmt"
	"net"
	"os"
)

func Put(conn net.Conn, request models.Request) {
	if request.Path[0:7] == "/files/" {
		filepath := request.Path[7:]
		dir := os.Args[2]
		writeData := request.Body
		readData, err := os.ReadFile(dir + "/" + filepath)
		if err != nil {
			err = utils.WriteResponse(conn, 404, nil, "")
			if err != nil {
				fmt.Println("Error writing response:", err.Error())
				return
			}
		} else {
			buffer := new(bytes.Buffer)
			buffer.Write(readData)
			buffer.Write([]byte(writeData + "\n"))
			data := buffer.Bytes()
			err = os.WriteFile(dir+"/"+filepath, data, 0644)
			if err != nil {
				fmt.Println("Error with writing to file :", err.Error())
				return
			}
			err = utils.WriteResponse(conn, 200, nil, "")
			if err != nil {
				fmt.Println("Error writing response:", err.Error())
				return
			}
		}
	}
}
