featurePath = $(PWD)
PKGS := $(shell go list ./... | grep -v /vendor)

fmt:
	find ! -path "./vendor/*" -name "*.go" -exec gofmt -s -w {} \;

gometalinter:
	gometalinter -D gotype -D aligncheck --vendor --deadline=600s --dupl-threshold=200 -e '_string' -j 5 ./...

doc-hunt:
	doc-hunt check -e

run-tests:
	./test.sh

run-quick-tests: setup-test-fixtures
	go test -v $(PKGS)

test-package:
	go test -race -cover -coverprofile=/tmp/chyle github.com/antham/chyle/$(pkg)
	go tool cover -html=/tmp/chyle -o /tmp/chyle.html
