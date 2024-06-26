package handlers

import (
	"codecrafters-http-server-go/app/models"
	"codecrafters-http-server-go/app/utils"
	"fmt"
	"net"
	"os"
)

func Delete(conn net.Conn, request models.Request) {
	if request.Path[0:7] == "/files/" {
		filepath := request.Path[7:]
		dir := os.Args[2]
		_, err := os.Stat(dir + "/" + filepath)
		if os.IsNotExist(err) {
			err = utils.WriteResponse(conn, 404, nil, "", request)
			if err != nil {
				fmt.Println("File not Found", err.Error())
			}
			return
		}
		err = os.Remove(dir + "/" + filepath)
		if err != nil {
			err = utils.WriteResponse(conn, 500, nil, "", request)
			fmt.Println("Error with deleting the file", err.Error())
			return
		}
		err = utils.WriteResponse(conn, 204, nil, "", request)
		if err != nil {
			fmt.Println("Error writing response:", err.Error())
			return
		}

	}
}
