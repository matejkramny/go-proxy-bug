.PHONY: all build dev server

all: build
	./bin/proxy-bug

build:
	go build -o ./bin/proxy-bug ./cli/proxy-bug

dev:
	@reflex -d none -sr '\.go$$' make server

server: build
	./bin/proxy-bug server
