package router

import (
	"app/codecrafters-http-server-go/handlers"
	"app/codecrafters-http-server-go/models"
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
