GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64

run: 
	go run cmd/server/server.go
