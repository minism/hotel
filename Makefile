SRCS=$(wildcard src/*.go)
TARGET=hotel

${TARGET}: ${SRCS}
	go build -o $@ $^

run: ${SRCS}
	go run $^

deps:
	dep ensure

docker-image: deps ${TARGET}
	docker build -t hotel .

docker-image-test: docker-image
	docker run --rm hotel

integration-test:
	python integration_test/__init__.py

clean:
	rm ${TARGET}
