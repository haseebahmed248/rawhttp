// Actual route login (health, echo, static)
package handler

import (
	"log"
	"os"
	"rawhttp/internal/request"
	"rawhttp/internal/response"
	"strings"
)

var mimeTypes = map[string]string{
	"html": "text/html",
	"css":  "text/css",
	"js":   "application/javascript",
	"json": "application/json",
	"png":  "image/png",
	"jpg":  "image/jpeg",
}

func HelloHandler(req *request.Request) *response.Response {
	return &response.Response{
		StatusCode: 200,
		Body:       []byte("Hello World"),
	}
}

func HealthHandler(req *request.Request) *response.Response {
	return &response.Response{
		StatusCode: 200,
		Body:       []byte("Healthy"),
	}
}

func EchoHandler(req *request.Request) *response.Response {
	return &response.Response{
		StatusCode: 200,
		Body:       []byte("Echo"),
	}
}

func StaticHandler(req *request.Request) *response.Response {
	var path = "./static"
	root, err := os.OpenRoot(path)
	if err != nil {
		log.Print(err)
		return &response.Response{
			StatusCode: 500,
			Body:       []byte("Internal Server Error"),
		}
	}
	defer root.Close()
	fileName := strings.SplitN(req.Path, "/", 3)
	data, err := root.ReadFile(fileName[2])
	if err != nil {
		return &response.Response{
			StatusCode: 404,
			Body:       []byte("File Not Found"),
		}
	}
	types := strings.SplitN(fileName[2], ".", 2)
	return &response.Response{
		StatusCode: 200,
		Body:       data,
		Headers: map[string]string{
			"Content-Type": mimeTypes[types[1]],
		},
	}
}
