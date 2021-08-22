FROM golang:1.16-alpine AS build
WORKDIR /src
RUN go get github.com/cespare/reflex
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/todo ./cmd/http-server

FROM alpine:3.13
RUN apk --no-cache -U upgrade
WORKDIR /app
COPY --from=build /app/todo .
USER guest
CMD ["/app/todo"]
