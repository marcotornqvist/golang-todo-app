FROM golang:1.14.6-alpine3.12 as builder

COPY go.mod go.sum /go/src/github.com/marcotornqvist/go-todo-app/

WORKDIR /go/src/github.com/marcotornqvist/go-todo-app

RUN go mod download

COPY . /go/src/github.com/marcotornqvist/go-todo-app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/go-todo-app github.com/marcotornqvist/go-todo-app

FROM alpine

COPY --from=builder /go/src/github.com/marcotornqvist/go-todo-app/.env .

RUN apk add --no-cache ca-certificates && update-ca-certificates

COPY --from=builder /go/src/github.com/marcotornqvist/go-todo-app/build/go-todo-app /usr/bin/go-todo-app

EXPOSE 8080 8080

ENTRYPOINT ["/usr/bin/go-todo-app"]