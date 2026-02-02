build:
	go build -o bin/app ./cmd

run: build
	./bin/app

test:
	go test ./...

fmt:
	go fmt ./...

doc:
	@echo "=== Документация по cmd ==="
	go doc ./cmd
	@echo "=== Документация по internal и всем подпакетам ==="
	@go list ./internal/... | xargs -r -n1 go doc

clean:
	rm -rf bin/

.PHONY: build run test fmt doc clean