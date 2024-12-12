LOCAL_BIN:=$(CURDIR)/bin
PROTO_PATH:=api/pb
PROTO_OUT_PATH:=pkg
PROTO_NAME:=user
PROTO_FILE:=${PROTO_PATH}/${PROTO_NAME}_v1/${PROTO_NAME}.proto

lint:
	${LOCAL_BIN}/golangci-lint run ./... --config .golangci.pipeline.yaml

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.3
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.35.1
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate-api:
	make generate-user-api

generate-user-api:
	@echo "Generate User API from proto files..."
	@mkdir -p ${PROTO_OUT_PATH}
	@protoc --proto_path ${PROTO_PATH} \
	--go_out=${PROTO_OUT_PATH} --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=${LOCAL_BIN}/protoc-gen-go \
	--go-grpc_out=${PROTO_OUT_PATH} --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=${LOCAL_BIN}/protoc-gen-go-grpc \
	${PROTO_FILE}
	@echo "Done!"