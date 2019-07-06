SRCS=$(wildcard src/*.go)
SERVER=server

${SERVER}: ${SRCS}
	go build -o $@ $^

run: ${SRCS}
	go run $^

watch: ${SRCS}
	cd src && gin run *.go

deps:
	dep ensure

integration-test:
	python integration_test/__init__.py
