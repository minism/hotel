SRCS=$(wildcard src/*/*.go)
MASTER_MAIN=services/master/main.go
MASTER_TARGET=hotel-master
SPAWNER_MAIN=services/spawner/main.go
SPAWNER_TARGET=hotel-spawner

${MASTER_TARGET}: ${MASTER_MAIN} ${SRCS} deps
	go build -o $@ $<

${SPAWNER_TARGET}: ${SPAWNER_MAIN} ${SRCS} deps
	go build -o $@ $<

run-master: ${MASTER_MAIN} ${SRCS} deps
	go run $<

run-spawner: ${SPAWNER_MAIN} ${SRCS} deps
	go run $<

deps:
	dep ensure

docker-image:
	docker build -t hotel .

integration-test:
	python test/integration_test.py

clean:
	rm -f ${MASTER_TARGET}
	rm -f ${SPAWNER_TARGET}
	rm -f data.db
