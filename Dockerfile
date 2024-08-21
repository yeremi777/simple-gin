FROM golang:1.23-alpine

WORKDIR /source-app

COPY . ./
RUN go mod download

RUN go build -o /app/

WORKDIR /

RUN rm -rf /source-app

ARG GO_APP_PORT=9999
ARG GO_DATABASE_HOST=127.0.0.1
ARG GO_DATABASE_PORT=5432
ARG GO_DATABASE_USERNAME=postgres
ARG GO_DATABASE_PASSWORD=root
ARG GO_DATABASE_NAME=simple_gin

ENV GO_APP_PORT=${PORT}
ENV GO_DATABASE_HOST=${DB_HOST}
ENV GO_DATABASE_PORT=${DB_PORT}
ENV GO_DATABASE_USERNAME=${DB_USERNAME}
ENV GO_DATABASE_PASSWORD=${DB_PASSWORD}
ENV GO_DATABASE_NAME=${DB_DATABASE}

EXPOSE 9999

CMD ["/app/simple-gin"]