
build:

	go clean && go build -o cici  -ldflags="-X main.BuildTime=`date +%Y-%m-%d.%H:%M:%S`" cici-backend.go