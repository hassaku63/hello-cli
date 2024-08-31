BINARY_DIR := bin

.PHONY: build
build:
	go build -o ${BINARY_DIR}/hello ./cmd/hello

.PHONY: clean
clean:
	rm -rf ${BINARY_DIR}/
