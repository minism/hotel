SRCS=$(wildcard src/*.go)
TARGET=hotel

${TARGET}: ${SRCS}
	go build -o $@ $^

run: ${SRCS}
	go run $^

docker-image: deps ${TARGET}
	docker build -t hotel .

watch: ${SRCS}
	cd src && gin -p 3001 run main.go

deps:
	go get -d -v ./...

# I'm not sure what the difference is here...
develop-deps:
	dep ensure

integration-test:
	python integration_test/__init__.py
