CURRENT_DATE=$(shell date +%Y-%m-%dT%H:%M:%S%z)
COVERAGE_FILE=./coverage.txt

.PHONY: test
test:
	@go test ./... -v

race:
	@go test -race ./... -v

.PHONY: coverage
coverage:
	@go test ./... -v -coverprofile=$(COVERAGE_FILE) -coverpkg=./...

.PHONY: show-coverage
show-coverage:
	@go tool cover -html=$(COVERAGE_FILE)

.PHONY: fmt
fmt:
	@gofmt -l .
