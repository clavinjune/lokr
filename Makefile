include tools.mk

setup:
	# try to use latest golang version
	@go install $(BUF)
	@go install $(wire)
	@go install $(mockery)

fmt:
	@gofmt -w -s .
	@go vet ./...

generate:
	@rm -rf ./api

mock:
	@mockery --all --with-expecter --output ./mocks

test:
	@go test -v -covermode=count -shuffle=on ./...

test/report:
	@go test -covermode=count -shuffle=on -coverprofile test-coverage.out -json ./... > test-report.json

test/cover: test/report
	@go tool cover -html=test-coverage.out

wire:
	@wire ./...