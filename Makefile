GOPATH	:= $(shell go env GOPATH)
GO		:= $(shell which go)

.PHONY: run
run: generate

.PHONY: test
test: generate
	$(GO) test -v ./...

.PHONY: generate
generate: hsm/hsm.proto
	PATH=${PATH}:"${GOPATH}/bin" protoc --go_out=pkg --go_opt=paths=source_relative \
    	--go-grpc_out=pkg --go-grpc_opt=paths=source_relative \
    	$<