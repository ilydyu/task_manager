FROM golang:1.26.2-alpine3.23 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /task_manager ./cmd

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

FROM alpine:3.23 AS run

COPY --from=build /go/bin/goose /usr/local/bin/goose
COPY --from=build /task_manager /task_manager
COPY --from=build /app/.env /.env
COPY migrations /migrations

EXPOSE 8080

CMD ["/task_manager"]