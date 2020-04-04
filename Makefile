## Run Lint
.PHONY: lint
lint:
	gofmt -s -w .
	goimports -w .
	go vet ./...
	golint -set_exit_status ./...

## Run Test
.PHONY: test
test: lint
	go test ./... -cover
