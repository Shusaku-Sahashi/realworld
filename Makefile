#/bin/sh

binfile=./bin

.PHONY: build
build:
	# TODO: @が必要な理由を調べる。(@がないと動かない。)
	@if [ -e ${binfile} ]; then \
		rm -fr ${binfile} & makdir ${binfile}; \
	fi 
	go build -o ${binfile}/app ./

.PHONY: run
run:
	go run ./
