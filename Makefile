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

run:
	life $(iterations)

clean:
	go clean ./...
