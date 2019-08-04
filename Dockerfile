# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# This is an updated version of the docker file which assumes you are able to
# build the binary artifact locally. Previous versions were dev environment
# free and could be deployed from anywhere. The problem is simply that it
# massively bloats the images, whereas this version *only* contains the
# binary artifact. I'm still not sure what is the best practice to use.

# Use iron which is a tiny microcontainer image for production deployments.
FROM iron/go

# Tag the maintainer.
MAINTAINER joshbothun@gmail.com

# Set the Current Working Directory inside the container.
WORKDIR /app

# Copy the binary to the working directory.
COPY ./hotel .

# The server runs on port 3000, so expose that.
EXPOSE 3000

# Setup a volume.
# VOLUME /data

# Run the executable.
ENTRYPOINT ["./hotel"]
