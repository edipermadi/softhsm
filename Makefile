GOPATH	:= $(shell go env GOPATH)
GO		:= $(shell which go)
PYTHON3	:= $(shell which python3)

.PHONY: run
run: generate

.PHONY: test
test: generate
	$(GO) test -v ./...

.PHONY: generate
generate: pb/pkcs11.proto
	PATH=${PATH}:"${GOPATH}/bin" protoc --go_out=pkg/transport --go_opt=paths=source_relative \
    	--go-grpc_out=pkg/transport --go-grpc_opt=paths=source_relative \
    	$<

.PHONY: flash-maixbit
flash-maixbit: kflash
	PATH=$(PATH):$(PWD)/kflash/bin tinygo flash -target=maixbit ./cmd/token

.PHONY: build-maixbit
build-maixbit: kflash
	PATH=$(PATH):$(PWD)/kflash/bin tinygo build -target=maixbit ./cmd/token

kflash:
	$(PYTHON3) -m venv kflash
	. kflash/bin/activate && pip install kflash