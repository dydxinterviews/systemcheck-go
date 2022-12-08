lint:
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint run --timeout=10m

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd

build: go.sum
ifeq ($(OS),Windows_NT)
	exit 1
else
	go build -mod=readonly $(BUILD_FLAGS) -o build ./cmd
endif

###############################################################################
###                           Tests & Simulation                            ###
###############################################################################

test: test-unit

test-unit:
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock norace' ./...
