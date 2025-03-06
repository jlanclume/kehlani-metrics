#!/bin/bash

SCRIPT_DIR=$(dirname "$(realpath "$0")")
TARGET_DIR="/target"

rm -rf "$TARGET_DIR/*"
 
# Find all .proto files and compile them
find "$SCRIPT_DIR/../messages" -name "*.proto" | while read file; do
  echo "Compiling $file..."
  protoc --proto_path="$SCRIPT_DIR/../messages" \
         --go_out="$TARGET_DIR" --go_opt=paths=source_relative \
         --go-grpc_out="$TARGET_DIR" --go-grpc_opt=paths=source_relative \
         "$file"
done

echo "Protobuf files have been compiled."
