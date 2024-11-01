# Variables
PROTO_DIR := ./src/protobuf
OUT_DIR := ./src/generated
PROTO_FILES := $(PROTO_DIR)/filemanager.proto

# Path to protoc (ensure it's installed and in your PATH)
PROTOC := protoc

# Target to generate the gRPC code from .proto files
generate:
	@echo "Generating gRPC code..."
	mkdir -p $(OUT_DIR)
	$(PROTOC) --go_out=$(OUT_DIR) --go_opt=paths=source_relative --go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative $(PROTO_FILES)

# Target to generate the gRPC code and run the Go server
run: generate
	@echo "Running Go server..."
	go run src/main.go

# Clean generated files (optional)
clean:
	@echo "Cleaning generated files..."
	rm -rf $(OUT_DIR)/*.pb.go

# Help target
help:
	@echo "Makefile commands:"
	@echo "  generate      - Generate gRPC code from .proto files"
	@echo "  run           - Generate gRPC code and run the Go server"
	@echo "  clean         - Clean generated .pb.go files"
