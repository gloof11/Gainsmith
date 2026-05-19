BINARY      := gainsmith
CMD         := ./main.go
BIN_DIR     := bin
BUILD_FLAGS ?= CGO_ENABLED=0 GOOS=linux
LD_FLAGS ?= -ldflags="-s -w"

.PHONY: run build clean test

run:
	go run $(CMD) serve

build:
	mkdir -p $(BIN_DIR)
	$(BUILD_FLAGS) go build -o $(BIN_DIR)/$(BINARY) $(LD_FLAGS) $(CMD) 

test:
	rm -rf test_pb_data
	cp -r pb_data test_pb_data 
	go test ./... -v
	rm -rf test_pb_data

test-cover:
	rm -rf test_pb_data
	cp -r pb_data test_pb_data 
	go test ./... -coverprofile cover.out -covermode count -v
	rm -rf test_pb_data

clean:
	rm -rf $(BIN_DIR)
	rm -rf test_pb_data
	rm -rf cover.out
