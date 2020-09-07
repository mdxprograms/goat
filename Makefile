.PHONY: build package rice

build:
	@CGO_ENABLED=0 go build -ldflags "-w" -a -o goat .

package: build
	${GOPATH}/bin/rice append -i . --exec goat
