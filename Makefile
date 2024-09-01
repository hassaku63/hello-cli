BINARY_DIR := bin

.PHONY: build
build:
	go build -o ${BINARY_DIR}/hello ./cmd/hello
	go build -o ${BINARY_DIR}/subcmd ./cmd/subcmd
	go build -o ${BINARY_DIR}/choice ./cmd/choice

.PHONY: clean
clean:
	rm -rf ${BINARY_DIR}/
