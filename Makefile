#/bin/sh

binfile=./bin

.PHONY: build
build: generate
	# TODO: @が必要な理由を調べる。(@がないと動かない。)
	@if [ -e ${binfile} ]; then \
		rm -fr ${binfile} && mkdir ${binfile} ; \
	fi 
	go build -o ./${binfile}/app ./

.PHONY: run
run: generate
	go run ./

.PHONY: lint
lint:
	golint ./...

.PHONY: fmt
fmt:
	go fmt ./

.PHONY: list 
list:
	go list -m all

.PHONY: generate
generate:
	go generate ./...

.PHONY: test
test:
	go test -cover ./...