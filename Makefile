.PHONY: build package start_api start_client dev

build: build_client
	@CGO_ENABLED=0 go build -ldflags "-w" -a -o goat .

package: build
	${GOPATH}/bin/rice append -i . --exec goat

start_api:
	go run main.go

start_client:
	cd client && npm start

build_client:
	cd client && npm run build

dev:
	make start_api & make start_client
