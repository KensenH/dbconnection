FROM golang:alpine3.14 AS build

WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/demo

ENTRYPOINT ["/bin/demo"]