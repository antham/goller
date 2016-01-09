compile:
	gox -output "build/{{.Dir}}_{{.OS}}_{{.Arch}}"
test:
	go test -cover ./...
vet:
	go vet ./...
check: test vet
fmt:
	gofmt -l -s -w .
