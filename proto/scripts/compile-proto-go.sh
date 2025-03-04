#!/bin/bash

mkdir -p /proto/generated/
rm -r /proto/generated/*

find /proto/messages -name "*.proto" | while read file; do
  echo "Compiling $file..."
  protoc --proto_path /proto/messages \
         --go_out=/proto/generated --go_opt=paths=source_relative \
         --go-grpc_out=/proto/generated --go-grpc_opt=paths=source_relative \
         "$file"
done

echo "Protobuf files have been compiled."