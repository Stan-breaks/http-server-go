package router

import (
	"codecrafters-http-server-go/app/handlers"
	"codecrafters-http-server-go/app/models"
	"net"
)

func Route(conn net.Conn, request models.Request) {
	switch request.Method {
	case "GET":
		handlers.Get(conn, request)
	case "POST":
		handlers.Post(conn, request)
	case "PUT":
		handlers.Put(conn, request)
	case "DELETE":
		handlers.Delete(conn, request)
	default:
		break
	}
}
