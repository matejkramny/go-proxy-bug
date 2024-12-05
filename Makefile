.PHONY: all build dev server proxy client

all: build
	./bin/proxy-bug

build:
	go build -o ./bin/proxy-bug ./cli/proxy-bug

dev:
	@reflex -d none -sr '\.go$$' make server

server: build
	./bin/proxy-bug server

proxy: build
	./bin/proxy-bug proxy

client: build
	./bin/proxy-bug client