SRCS=$(wildcard src/*/*.go)
PROTO_SRCS=$(wildcard proto/*.proto)
PROTO_OUTS=$(patsubst proto/%.proto, src/proto/%.pb.go, $(PROTO_SRCS))
MASTER_MAIN=services/master/main.go
MASTER_TARGET=hotel-master
SPAWNER_MAIN=services/spawner/main.go
SPAWNER_TARGET=hotel-spawner

${MASTER_TARGET}: ${MASTER_MAIN} ${SRCS} ${PROTO_OUTS}
	go build -o $@ $<

${SPAWNER_TARGET}: ${SPAWNER_MAIN} ${SRCS} ${PROTO_OUTS}
	go build -o $@ $<

${PROTO_OUTS}: ${PROTO_SRCS}
	mkdir -p src/proto
	protoc --go_out=plugins=grpc:src $^

run-master: ${MASTER_MAIN} ${SRCS} ${PROTO_OUTS}
	go run $<

run-spawner: ${SPAWNER_MAIN} ${SRCS} ${PROTO_OUTS}
	go run $<

deps:
	dep ensure

protos: ${PROTO_OUTS}

docker-images:
	docker-compose build

integration-test:
	python test/integration_test.py

clean:
	rm -f ${MASTER_TARGET}
	rm -f ${SPAWNER_TARGET}
	rm -f data.db
