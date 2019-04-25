SRCS=$(wildcard src/*.go)
SERVER=server

${SERVER}: ${SRCS}
	go build -o $@ $^ 
