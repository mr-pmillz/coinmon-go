BIN="./bin"
DIST="./dist"
SRC=$(shell find . -name "*.go")
CURRENT_TAG=$(shell git describe --tags --abbrev=0)

GOLANGCI := $(shell command -v golangci-lint 2>/dev/null)
RICHGO := $(shell command -v richgo 2>/dev/null)
GOTESTFMT := $(shell command -v gotestfmt 2>/dev/null)
MIN_GOLANGCI_LINT_VERSION := 001043000

.PHONY: fmt lint build test clean compile compress

default: all

all: fmt lint build test release

release: clean build compile compress

fmt:
	$(info ******************** checking formatting ********************)
	@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)

.PHONY: golangci-lint-check
golangci-lint-check:
	$(eval GOLANGCI_LINT_VERSION := $(shell printf "%03d%03d%03d" $(shell golangci-lint --version | grep -Eo '[0-9]+\.[0-9.]+' | tr '.' ' ');))
	$(eval MIN_GOLANGCI_LINT_VER_FMT := $(shell printf "%g.%g.%g" $(shell echo $(MIN_GOLANGCI_LINT_VERSION) | grep -o ...)))
	@hash golangci-lint > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		echo "Downloading golangci-lint v${MIN_GOLANGCI_LINT_VER_FMT}"; \
		export BINARY="golangci-lint"; \
		curl -sfL "https://raw.githubusercontent.com/golangci/golangci-lint/v${MIN_GOLANGCI_LINT_VER_FMT}/install.sh" | sh -s -- -b $(GOPATH)/bin v$(MIN_GOLANGCI_LINT_VER_FMT); \
	elif [ "$(GOLANGCI_LINT_VERSION)" -lt "$(MIN_GOLANGCI_LINT_VERSION)" ]; then \
		echo "Downloading newer version of golangci-lint v${MIN_GOLANGCI_LINT_VER_FMT}"; \
		export BINARY="golangci-lint"; \
		curl -sfL "https://raw.githubusercontent.com/golangci/golangci-lint/v${MIN_GOLANGCI_LINT_VER_FMT}/install.sh" | sh -s -- -b $(GOPATH)/bin v$(MIN_GOLANGCI_LINT_VER_FMT); \
	fi

.PHONY: lint
lint: golangci-lint-check
	$(info ******************** running lint tools ********************)
	golangci-lint run -c .golangci-lint.yml -v ./... --timeout 10m

test:
	$(info ******************** running tests ********************)
    ifeq ($(GITHUB_ACTIONS), true)
        ifndef GOTESTFMT
			$(warning "could not find gotestfmt in $(PATH), running: go install github.com/haveyoudebuggedit/gotestfmt/v2/cmd/gotestfmt@latest")
			$(shell go install github.com/haveyoudebuggedit/gotestfmt/v2/cmd/gotestfmt@latest)
        endif
		go test -json -v ./... 2>&1 | tee coverage/gotest.log | gotestfmt
    else
        ifndef RICHGO
			$(warning "could not find richgo in $(PATH), running: go get github.com/kyoh86/richgo")
			$(shell go get -u github.com/kyoh86/richgo)
        endif
		richgo test -v ./...
    endif

clean:
	rm -rf $(BIN) 2>/dev/null
	rm -rf $(DIST) 2>/dev/null

build:
	go env -w GOFLAGS=-mod=mod
	go mod tidy
	go build -v -trimpath -ldflags="-s -w" .

compile:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/coinmon-go-$(CURRENT_TAG)-linux-amd64 -trimpath -ldflags="-s -w" -gcflags="all=-trimpath=${GOPATH}" -asmflags="all=-trimpath=${GOPATH}" main.go
	GOOS=linux GOARCH=arm64 go build -o bin/linux/arm64/coinmon-go-$(CURRENT_TAG)-linux-arm64 -trimpath -ldflags="-s -w" -gcflags="all=-trimpath=${GOPATH}" -asmflags="all=-trimpath=${GOPATH}" main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/amd64/coinmon-go-$(CURRENT_TAG)-x86_64-apple-darwin_amd64 -trimpath -ldflags="-s -w" -gcflags="all=-trimpath=${GOPATH}" -asmflags="all=-trimpath=${GOPATH}" main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/darwin/arm64/coinmon-go-$(CURRENT_TAG)-x86_64-apple-darwin_arm64 -trimpath -ldflags="-s -w" -gcflags="all=-trimpath=${GOPATH}" -asmflags="all=-trimpath=${GOPATH}" main.go
	GOOS=windows GOARCH=amd64 go build -o bin/windows/amd64/coinmon-go-$(CURRENT_TAG)-x86_64-windows_amd64.exe -trimpath -ldflags="-s -w" -gcflags="all=-trimpath=${GOPATH}" -asmflags="all=-trimpath=${GOPATH}" main.go
	GOOS=windows GOARCH=386 go build -o bin/windows/386/coinmon-go-$(CURRENT_TAG)-386-windows.exe -trimpath -ldflags="-s -w" -gcflags="all=-trimpath=${GOPATH}" -asmflags="all=-trimpath=${GOPATH}" main.go

compress:
	gzip -9 bin/linux/amd64/coinmon-go-$(CURRENT_TAG)-linux-amd64
	gzip -9 bin/linux/arm64/coinmon-go-$(CURRENT_TAG)-linux-arm64
	gzip -9 bin/darwin/amd64/coinmon-go-$(CURRENT_TAG)-x86_64-apple-darwin_amd64
	gzip -9 bin/darwin/arm64/coinmon-go-$(CURRENT_TAG)-x86_64-apple-darwin_arm64
	gzip -9 bin/windows/386/coinmon-go-$(CURRENT_TAG)-386-windows.exe
	gzip -9 bin/windows/amd64/coinmon-go-$(CURRENT_TAG)-x86_64-windows_amd64.exe

