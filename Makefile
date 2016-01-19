compile:
	gox -output "build/{{.Dir}}_{{.OS}}_{{.Arch}}"
test:
	go test -race -cover ./...
test-unary:
	go test -race -cover -coverprofile=/tmp/goller github.com/antham/goller/$(pkg)
vet:
	go vet ./...
check: test vet fmt
fmt:
	gofmt -l -s -w .
