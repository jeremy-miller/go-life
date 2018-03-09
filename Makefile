default: test

setup: setup-go setup-dep

setup-go:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

setup-dep:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	dep ensure

.PHONY: test
test: lint
	go test -v

.PHONY: lint
lint:
	gometalinter --tests --vendor

build: lint
	go install ./cmd/...

# e.g. make life iterations=3
run:
	life $(iterations)

clean:
	go clean ./...

# e.g. make simplify files=internal/life/
.PHONY: simplify
simplify:
	gofmt -s -d $(files)
