// HTTP request parsing
package request

import (
	"bufio"
	"errors"
	"log"
	"strconv"
	"strings"
)

type Request struct {
	Method  string            // "GET"
	Path    string            // "/users"
	Query   string            // "id=123"
	Version string            // "HTTP/1.1"
	Headers map[string]string // {"Host": "localhost:9090", ...}
	Body    []byte
}

func Parse(reader *bufio.Reader) (*Request, error) {
	if reader == nil {
		return nil, errors.New("empty")
	}
	var method, path, version string
	var bodyLength int
	query := ""
	headers := make(map[string]string)

	status, err := reader.ReadString('\n')

	if err != nil {
		log.Print("Error: ", err)
		return nil, err
	}
	parts := strings.Split(status, " ")

	method = parts[0] // "GET"
	pathParts := strings.Split(parts[1], "?")
	path = pathParts[0] // "/hello"
	if len(pathParts) == 2 {
		query = pathParts[1]
	}
	version = strings.TrimSpace(parts[2]) // "HTTP/1.1"

	for {
		header, err := reader.ReadString('\n')
		if err != nil {
			log.Print("Error: ", err)
			return nil, err
		}
		if strings.TrimSpace(header) == "" {
			break
		}
		data := strings.SplitN(header, ":", 2)
		if data[0] == "Content-Length" {
			bodyLength, _ = strconv.Atoi(strings.TrimSpace(data[1]))
		}
		headers[data[0]] = strings.TrimSpace(data[1])
	}
	body := make([]byte, bodyLength)
	reader.Read(body)
	return &Request{
		Method:  method,
		Path:    path,
		Query:   query,
		Version: version,
		Headers: headers,
		Body:    body,
	}, nil
}
