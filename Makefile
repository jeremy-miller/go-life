PKG_LIST := $(shell go list ./... | grep -v /vendor/)

default: build

.PHONY: setup
setup:
	@curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	@go get -u github.com/alecthomas/gometalinter
	@go get -u github.com/m3db/build-tools/linters/badtime
	@go get -u github.com/hypnoglow/durcheck
	@gometalinter --install
	@dep ensure

.PHONY: dep
dep:
	@dep ensure

# check dependency versions
.PHONY: dep-versions
dep-versions:
	@dep status

.PHONY: lint
lint:
	@gometalinter --tests --vendor ./...

.PHONY: build
build:
	@go install $(PKG_LIST)

.PHONY: test
test:
	@go test -v -race $(PKG_LIST)

.PHONY: coverage
coverage:
	@go test $(PKG_LIST) -v -coverprofile .testCoverage.txt
	@go tool cover -func=.testCoverage.txt

.PHONY: run
# e.g. make life iterations=3
run:
	@life $(iterations)

.PHONY: clean
clean:
	@go clean $(PKG_LIST)
	@rm -rf ./.testCoverage.txt

.PHONY: simplify
simplify:
	@$(foreach file,$(PKG_LIST),gofmt -s -w $(GOPATH)/src/$(file);)
