.PHONY: build build_client package start_api start_client dev

APP_NAME="goat"

build: build_client
	@CGO_ENABLED=0 go build -ldflags "-w" -a -o $(APP_NAME) .

package: build
	${GOPATH}/bin/rice append -i . --exec $(APP_NAME)

start_api:
	go run main.go

start_client:
	cd client && npm start

build_client:
	cd client && npm run build

dev:
	make start_api & make start_client
