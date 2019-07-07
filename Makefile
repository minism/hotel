SRCS=$(wildcard src/*.go)
TARGET=hotel

${TARGET}: ${SRCS}
	go build -o $@ $^

run: ${SRCS}
	go run $^

watch: ${SRCS}
	cd src && gin -p 3001 run main.go

deps:
	dep ensure

integration-test:
	python integration_test/__init__.py
