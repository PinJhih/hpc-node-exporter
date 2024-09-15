BINARY_NAME=hpc-node-exporter
BUILD_DIR=build

.PHONY: all
all: build

build: $(BUILD_DIR)/$(BINARY_NAME)
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd

.PHONY: run
run:
	go run ./cmd/hpc-node-exporter.go

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)/*