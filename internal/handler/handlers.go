// Actual route login (health, echo, static)
package handler

import (
	"rawhttp/internal/request"
	"rawhttp/internal/response"
)

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
