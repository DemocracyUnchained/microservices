# Adapted from:
# https://github.com/docker-library/golang/blob/master/1.8/onbuild/Dockerfile
FROM golang:1.8

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

# Copy over the configuration script for the Go app.
COPY config.toml /etc/config.toml

# Make the directory where the code will be mounted as a volume.
RUN mkdir -p "$GOPATH/src/democracy"
VOLUME "$GOPATH/src/democracy"

# Copy startup/wrapper scripts to the container.
COPY go-wrapper /usr/local/bin/
COPY democracy.sh /usr/local/bin/
COPY wait-for-it.sh /usr/local/bin/

ENV PATH /usr/local/bin:$PATH
ENV PATH $GOPATH/src/democracy:$PATH
