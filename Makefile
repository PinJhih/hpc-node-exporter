BINARY_NAME=hpc-node-exporter
BUILD_DIR=build

.PHONY: all run run-bin clean
all: $(BUILD_DIR)/$(BINARY_NAME)

$(BUILD_DIR)/$(BINARY_NAME):
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd

run:
	go run ./cmd/hpc-node-exporter.go

run-bin:
	make all
	./$(BUILD_DIR)/$(BINARY_NAME)

clean:
	rm -rf $(BUILD_DIR)/*