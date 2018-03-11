PROJECT_DIR := "life-go"
PKG := "github.com/jeremy-miller/$(PROJECT_DIR)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

.PHONY: setup dep lint build test race run clean simplify

default: build

setup:
	@curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	@$(MAKE) dep

dep:
	@go get -u github.com/alecthomas/gometalinter
	@gometalinter --install
	@dep ensure

lint:
	@gometalinter --tests --vendor

build: lint
	@go install $(PKG_LIST)

test: lint
	@go test -v $(PKG_LIST)

race: lint
	@go test -v -race $(PKG_LIST)

# e.g. make life iterations=3
run:
	@life $(iterations)

clean:
	@go clean $(PKG_LIST)
	@rm -rf ./cmd/life/debug

# e.g. make simplify files=internal/life/
simplify:
	@gofmt -s -d $(files)
