.PHONY: build package

build:
	CGO_ENABLED=0 go build -ldflags "-w" -a -o goat .

package: build
	rice append -i . --exec goat
