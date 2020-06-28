
build-backend:
	cd cici-backend;go clean && GOOS=linux GOARCH=amd64 go build -o cici  -ldflags="-X main.BuildStamp=`date +%Y-%m-%d.%H:%M:%S`" cici-backend.go


build-frontend:
	cd cici-frontend; yarn run build


install-frontend-dep:
	cd cici-frontend; yarn install
