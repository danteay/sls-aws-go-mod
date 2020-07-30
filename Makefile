.PHONY: build clean test lint build_sam_template deploy

install:
	@go mod download
	@npm i # install all serverless plugin dependencies

build: clean
	@export GO111MODULE=on
	@sh build.sh

clean:
	@rm -rf ./bin

test:
	@go test -v -race `go list ./... | grep -v node_modules`

lint:
	@golangci-lint run ./...

build_sam_template:
	@sls sam export --output ./template.yml

run: build build_sam_template
	@sam local start-api

deploy: build
	@sls deploy --verbose
