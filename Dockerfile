FROM golang:1.25.4-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY . .

ENV CGO_ENABLED=1

RUN go build -o app

FROM alpine

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/templates ./templates

EXPOSE 8080

CMD ["./app"]
