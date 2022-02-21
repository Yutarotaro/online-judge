FROM golang:1.15.6-alpine3.12 as builder

RUN apk update && apk add --no-cache git curl

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /main .

FROM alpine:3.12
COPY --from=builder /main .
ENV PORT=${PORT}
ENTRYPOINT ["/main web"]