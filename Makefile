run dev:
	- go run ./cmd/web -addr=":8443"
run test:
	- go test -v ./...
help:
	- go run ./cmd/web -help
