FROM golang:1.24.0-alpine3.20 AS builder

WORKDIR /app

# Install build dependencies for CGO
RUN apk add --no-cache gcc musl-dev

# Set CGO_ENABLED=1
ENV CGO_ENABLED=1

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/url_shortener

FROM alpine:3.21.2

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]