build:
	go build -o bin/app ./cmd

run: build
	./bin/app

test:
	go test ./...

fmt:
	go fmt ./...

doc:
	godoc -http=:6060

clean:
	rm -rf bin/

.PHONY: build run test fmt doc clean