
.PHONY: build
build:
	go build -o build/airfog ./cmd/airfog

.PHONY: build-all
build-all:
	GOOS=darwin GOARCH=arm64 go build -o build/airfog-osx-arm ./cmd/airfog
	GOOS=linux GOARCH=386 go build -o build/airfog-linux ./cmd/airfog

