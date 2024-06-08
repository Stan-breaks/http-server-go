package utils

var HTTPStatusCodes = map[int]string{
	100: "100 Continue",
	101: "101 Switching Protocols",
	200: "200 OK",
	201: "201 Created",
	204: "204 No Content",
	301: "301 Moved Permanently",
	302: "302 Found",
	304: "304 Not Modified",
	400: "400 Bad Request",
	401: "401 Unauthorized",
	403: "403 Forbidden",
	404: "404 Not Found",
	405: "405 Method Not Allowed",
	500: "500 Internal Server Error",
	501: "501 Not Implemented",
	502: "502 Bad Gateway",
	503: "503 Service Unavailable",
}
