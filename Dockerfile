# Dockerfile References: https://docs.docker.com/engine/reference/builder/

########################################################################
# This docker file is structured as a multi-stage build. The first stage
# sets up an environment whose responsibility is simply to compile the
# go binary artifact with all needed dependencies.

FROM golang:alpine AS builder

# Tag the maintainer.
MAINTAINER joshbothun@gmail.com

# Requirements for compiling go-sqlite3.  This could be a base image.
# https://github.com/mattn/go-sqlite3
RUN apk add --update gcc musl-dev

# Set the working directory to where go will expect packages.
WORKDIR $GOPATH/src/minornine.com/hotel/

# Copy everything into the container (Except whats in .dockerignore).
COPY . .

# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
# Note that this step does nothing if "dep ensure" was run locally first, which is a nice optimization.
RUN go get -d -v ./...

# Build the artifact.
RUN go build -o /build/hotel src/*.go

#####################################################################
# The second stage produces the runtime image which is the minimal
# environment needed to execute the binary artifact.
# Note that its important the "build" and "runtime" linux environments
# are both alpine here, so we don't need to think about cross-compiling.
FROM alpine

# Set the Current Working Directory inside the container.
WORKDIR /app

# Copy the artifact from the first stage.
COPY --from=builder /build/hotel  .

# Copy the default config to the working directory.
COPY example.config.json .

# The server runs on port 3000, so expose that.
EXPOSE 3000

# Setup a volume.
VOLUME /data

# Run the executable.
ENTRYPOINT ["./hotel", "--data-path=/data"]
