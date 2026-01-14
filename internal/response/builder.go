// HTTP response construction
package response

import (
	"fmt"
	"net"
)

type Response struct {
	StatusCode int
	StatusText string // "OK", "Not Found", "Internal Server Error"
	Headers    map[string]string
	Body       []byte
}

func (r *Response) Write(conn net.Conn) error {
	var data string
	data = fmt.Sprintf("HTTP/1.1 %d %s\r\n", r.StatusCode, r.StatusText)
	for k, v := range r.Headers {
		data += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	data += "Content-Length: " + fmt.Sprintf("%d\r\n", len(r.Body))
	data += "\r\n"

	data += string(r.Body)
	_, err := conn.Write([]byte(data))
	if err != nil {
		return err
	}
	return nil
}
