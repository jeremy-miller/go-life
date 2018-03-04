# default: lint test

# .PHONY lint
# lint:


# .PHONY test
# test:


setup: setup-go setup-dep

setup-go:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

setup-dep:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	dep ensure
