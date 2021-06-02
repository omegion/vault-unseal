export PATH := $(abspath ./vendor/bin):$(PATH)

ifndef VERSION
   GIT_VERSION = $(shell git describe --tags --always 2> /dev/null || echo 0.0.0)
else
   GIT_VERSION = $(VERSION)
endif

BASE_PACKAGE_NAME  = github.com/omegion/vault-unseal
LDFLAGS            = -ldflags "-X $(BASE_PACKAGE_NAME)/internal/info.Version=$(GIT_VERSION)"
BUFFER            := $(shell mktemp)
REPORT_DIR         = dist/report
COVER_PROFILE      = $(REPORT_DIR)/coverage.out

.PHONY: build
build:
	CGO_ENABLED=0 go build $(LDFLAGS) -installsuffix cgo -o dist/vault-unseal main.go

build-for-container:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -a -installsuffix cgo -o dist/vault-unseal-linux main.go

.PHONY: lint
lint:
	@echo "Checking code style"
	gofmt -l . | tee $(BUFFER)
	@! test -s $(BUFFER)
	go vet ./...
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.1
	@golangci-lint --version
	golangci-lint run --fix
	go get -u golang.org/x/lint/golint
	golint -set_exit_status ./...

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
	docker build -t vault-unseal .
	docker tag vault-unseal:latest omegion/vault-unseal:$(GIT_VERSION)
	docker push omegion/vault-unseal:$(GIT_VERSION)
