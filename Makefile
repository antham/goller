compile:
	gox -output "build/{{.Dir}}_{{.OS}}_{{.Arch}}"
