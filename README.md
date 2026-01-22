# rawhttp

A custom HTTP/1.1 server built from raw TCP sockets in Go — without using `net/http`.

## Goal

Understand HTTP at the protocol level by implementing:
- TCP connection handling
- HTTP request parsing
- Response building
- Routing
- Static file serving
- JSON API support

## Run

```bash
go run ./cmd/server/main.go 9090
```

Test with:
```bash
curl http://localhost:9090/
```

## Why

Built to learn networking fundamentals, not to replace production servers.
