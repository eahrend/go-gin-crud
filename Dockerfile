# syntax=docker/dockerfile:experimental

# Build Image
ARG GO_VERSION=1.17
FROM golang:${GO_VERSION}-alpine AS builder
RUN apk add --no-cache --update \
        openssh-client \
        ca-certificates \
        build-base
RUN mkdir -p /go/src/github.com/eahrend/go-gin-crud
WORKDIR /go/src/github.com/eahrend/go-gin-crud
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app .

# Application layer
FROM alpine
ARG PORT
ENV PORT 8080
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates
RUN apk add bash
RUN mkdir -p /app/migrations
COPY --from=builder /app /app
COPY ./db/migrations/ /app/migrations
COPY ./wait-for-it.sh /app/
RUN chmod +x /app/wait-for-it.sh
EXPOSE $PORT
RUN adduser -D appuser && chown -R appuser /app
USER appuser
WORKDIR /app
CMD ["./app"]