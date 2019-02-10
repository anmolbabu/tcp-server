PROJECT := github.com/anmolbabu/tcp-server
GITCOMMIT := $(shell git rev-parse --short HEAD 2>/dev/null)
PKGS := $(shell go list  ./... | grep -v $(PROJECT)/vendor)
BUILD_FLAGS := -ldflags="-w -X $(PROJECT)/cmd.GITCOMMIT=$(GITCOMMIT)"
FILES := tcp-server dist
HOME := $(shell echo $(HOME))
default: bin

.PHONY: bin
bin:
	go build ${BUILD_FLAGS} -o tcp-server main.go

.PHONY: run
run: bin
	mkdir -p $(HOME)/.tcp-server/
	cp ./tcp-server.conf $(HOME)/.tcp-server/tcp-server.conf
	./tcp-server