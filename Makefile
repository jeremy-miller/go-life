default: build

.PHONY: setup
setup: setup-go setup-dep

.PHONY: setup-go
setup-go:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

.PHONY: setup-dep
setup-dep:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	dep ensure

.PHONY: test
test: lint
	go test -v ./...

.PHONY: lint
lint:
	gometalinter --tests --vendor

.PHONY: build
build: lint
	go install ./cmd/...

.PHONY: run
# e.g. make life iterations=3
run:
	life $(iterations)

.PHONY: clean
clean:
	go clean ./...
	rm -rf ./cmd/life/debug

# e.g. make simplify files=internal/life/
.PHONY: simplify
simplify:
	gofmt -s -d $(files)
