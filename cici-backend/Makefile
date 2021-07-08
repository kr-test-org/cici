
build-backend:
	go clean && GOOS=linux GOARCH=amd64 go build -o cici  -ldflags="-X main.BuildStamp=`date +%Y-%m-%d.%H:%M:%S`" cici-backend.go
