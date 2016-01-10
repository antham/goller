compile:
	gox -output "build/{{.Dir}}_{{.OS}}_{{.Arch}}"
test:
	go test -race -cover ./...
vet:
	go vet ./...
check: test vet fmt
fmt:
	gofmt -l -s -w .
