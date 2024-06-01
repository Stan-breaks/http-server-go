# Simple HTTP Server

This is a simple HTTP server implemented in Go. It supports the following functionalities:

## Features

1. Echo Server: Upon receiving a GET request to /echo/..., the server will respond with the text following /echo/ as the response body.
2. Compression Support: The server can compress the response body using gzip encoding if the client requests it via the Accept-Encoding header.
3. File Serving: The server can serve files from a specified directory by sending a GET request to /files/..., where ... represents the file path relative to the specified directory.
4. File Upload: The server can handle file uploads by accepting PUT requests to /files/..., where ... represents the file path relative to the specified directory.
5. User-Agent Header: The server can echo back the User-Agent header value by sending a GET request to /user-agent.

Usage

Compile the Go code and run it:

```bash
go run app/server.go
```
