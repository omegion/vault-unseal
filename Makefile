export PATH := $(abspath ./vendor/bin):$(PATH)

ifndef VERSION
   GIT_VERSION = $(shell git describe --tags --always 2> /dev/null || echo 0.0.0)
else
   GIT_VERSION = $(VERSION)
endif

BASE_PACKAGE_NAME  = github.com/omegion/vault-unseal
LDFLAGS            = -ldflags "-buildid=$(GIT_VERSION)"
BUFFER            := $(shell mktemp)
REPORT_DIR         = dist/report
COVER_PROFILE      = $(REPORT_DIR)/coverage.out
TARGETOS		   = darwin
TARGETARCH		   = amd64
BINARY_NAME        = vault-unseal
BINARY_PATH        = dist/$(BINARY_NAME)

.PHONY: build
build:
	CGO_ENABLED=0 GOOS="$(TARGETOS)" GOARCH="$(TARGETARCH)" go build $(LDFLAGS) -a -installsuffix cgo -o $(BINARY_PATH) main.go

.PHONY: lint
lint:
	@echo "Checking code style"
	gofmt -l . | tee $(BUFFER)
	@! test -s $(BUFFER)
	go vet ./...

	# golangci-lint
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.45.2
	@golangci-lint --version
	golangci-lint run

	# Statuscheck
	go install honnef.co/go/tools/cmd/staticcheck@2022.1
	staticcheck ./...

.PHONY: test
test:
	@echo "Running unit tests"
	mkdir -p $(REPORT_DIR)
	go test -covermode=count -coverprofile=$(COVER_PROFILE) -tags test -failfast ./...
	go tool cover -html=$(COVER_PROFILE) -o $(REPORT_DIR)/coverage.html

.PHONY: cut-tag
cut-tag:
	@echo "Cutting $(version)"
	git tag $(version)
	git push origin $(version)

.PHONY: release
release: build-for-container
	@echo "Releasing $(GIT_VERSION)"
	docker build -t $(BINARY_NAME) .
	docker tag $(BINARY_NAME):latest omegion/$(BINARY_NAME):$(GIT_VERSION)
	docker push omegion/$(BINARY_NAME):$(GIT_VERSION)
