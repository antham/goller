compile:
	gox -output "build/{{.Dir}}_{{.OS}}_{{.Arch}}"
test-unary:
	go test -race -cover github.com/antham/goller/$(pkg)
test-all:
	go test -race ./...
coverage-profile:
	go test -cover -coverprofile=/tmp/goller github.com/antham/goller/$(pkg)
	go tool cover -html=/tmp/goller
coverage-all:
	./test.sh
vet:
	go vet ./...
run-tests: coverage-all vet fmt
fmt:
	gofmt -l -s -w .
version:
	git stash -u
	sed -i "s/[[:digit:]]\+\.[[:digit:]]\+\.[[:digit:]]\+/$(v)/g" main.go
	git add -A
	git commit -m "feat(version) : "$(v)
	git tag v$(v) master
