CURRENT_DATE=$(shell date +%Y-%m-%dT%H:%M:%S%z)
COVERAGE_FILE=./coverage.txt

.PHONY: test
test:
	@go test ./tests/... -v

race:
	@go test -race ./tests/... -v

.PHONY: coverage
coverage:
	@go test ./tests/... -v -coverprofile=$(COVERAGE_FILE) -coverpkg=./...

.PHONY: show-coverage
show-coverage:
	@go tool cover -html=$(COVERAGE_FILE)

.PHONY: fmt
fmt:
	@gofmt -l .
