PROTOC = $(shell which protoc)
PROTOC_GEN_GO = $(shell which protoc-gen-go)
NAIS_API_COMMIT_SHA := bfae2d608d86cb6ea9e3a721ad86b2bfe0c1c942
NAIS_API_TARGET_DIR=internal/naisapi/protoapi

.PHONY: all

all: generate test check fmt build

generate: generate-nais-api

build:
	go build -o bin/garbage ./cmd/garbage

test:
	go test ./...

check: staticcheck vulncheck deadcode

staticcheck:
	go run honnef.co/go/tools/cmd/staticcheck@latest ./...

vulncheck:
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

deadcode:
	go run golang.org/x/tools/cmd/deadcode@latest -test ./...

fmt:
	go run mvdan.cc/gofumpt@latest -w ./

generate-nais-api:
	mkdir -p ./$(NAIS_API_TARGET_DIR)
	wget -O ./$(NAIS_API_TARGET_DIR)/cluster.proto https://raw.githubusercontent.com/nais/api/$(NAIS_API_COMMIT_SHA)/pkg/protoapi/schema/cluster.proto
	$(PROTOC) \
		--proto_path=$(NAIS_API_TARGET_DIR) \
		--go_opt=Mcluster.proto=github.com/nais/garbage/$(NAIS_API_TARGET_DIR) \
		--go_opt=paths=source_relative \
		--go_out=$(NAIS_API_TARGET_DIR) \
		--go-grpc_opt=Mcluster.proto=github.com/nais/garbage/$(NAIS_API_TARGET_DIR) \
		--go-grpc_opt=paths=source_relative \
		--go-grpc_out=$(NAIS_API_TARGET_DIR) \
		$(NAIS_API_TARGET_DIR)/*.proto
	rm -f $(NAIS_API_TARGET_DIR)/*.proto
