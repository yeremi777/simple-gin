FROM golang:1.23-alpine

WORKDIR / .

COPY . ./

RUN go mod tidy

ENTRYPOINT [ "/bin/sh", "startup.sh" ]