run dev:
	- go run ./cmd/web -addr=":8443"
run test:
	- go test -v ./...
run fasttest:
	- go test -v -short ./...
run coverage:
	- go test -v -cover ./...
help:
	- go run ./cmd/web -help
