FROM golang:1.24-alpine

# Install dependencies
RUN apk add --no-cache inotify-tools bash git protobuf protobuf-dev

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.5 && \
 go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1