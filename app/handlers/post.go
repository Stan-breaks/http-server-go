package handlers

import (
	"codecrafters-http-server-go/app/models"
	"codecrafters-http-server-go/app/utils"
	"fmt"
	"net"
	"os"
)

func Post(conn net.Conn, request models.Request) {
	if request.Path[0:7] == "/files/" {
		filepath := request.Path[7:]
		dir := os.Args[2]
		data := request.Body
		err := os.WriteFile(dir+"/"+filepath, []byte(data+"\n"), 0644)
		if err != nil {
			fmt.Println("Error writing into file", err.Error())
			return
		}
		err = utils.WriteResponse(conn, 201, nil, "",request)
		if err != nil {
			fmt.Println("Error writing response:", err.Error())
			return
		}

	}
}
