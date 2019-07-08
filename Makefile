all: lint install

build: go.sum
		GOOS=linux go build $(BUILD_FLAGS) -o build/linkd ./cmd/linkd
		GOOS=linux go build $(BUILD_FLAGS) -o build/linkcli ./cmd/linkcli

install: go.sum
		go install $(BUILD_FLAGS) ./cmd/linkd
		go install $(BUILD_FLAGS) ./cmd/linkcli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

lint:
	golangci-lint run
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s
	go mod verify

build-docker-linkdnode:
	$(MAKE) -C networks/local

localnet-init: build build-docker-linkdnode
	docker run --rm -v $(CURDIR)/build:/linkd:Z kfangw/linkdnode testnet --v 5 -o . --starting-ip-address 192.168.10.2

localnet-init-single: build build-docker-linkdnode
	docker run --rm -v $(CURDIR)/build:/linkd:Z kfangw/linkdnode testnet --v 1 -o . --starting-ip-address 192.168.10.2

.PHONY: all build install lint build-docker-linkdnode