#/bin/sh

binfile=./bin

.PHONY: build
build: ## go buildを実行する。
	# TODO: @が必要な理由を調べる。(@がないと動かない。)
	@if [ -e ${binfile} ]; then \
		rm -fr ${binfile} & makdir ${binfile}; \
	fi 
	go build -o ${binfile}/app ./

.PHONY: run
run:
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
