.PHONY: all

all:
	mkdir -p bin
	go build -o bin/server ./server
	go build -o bin/client ./client

clean:
	rm -rf bin
