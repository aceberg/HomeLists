mod:
	rm go.mod || true && \
	rm go.sum || true && \
	go mod init github.com/aceberg/HomeLists && \
	go mod tidy

run:
	cd cmd/HomeLists && \
	go run .

fmt:
	go fmt ./...

lint:
	golangci-lint run
	golint ./...

check: fmt lint