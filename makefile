test:
	@go test ./... -coverprofile=coverage.out -coverpkg=./...
html-coverage: test
	@go tool cover -html=coverage.out -o coverage.html && open coverage.html