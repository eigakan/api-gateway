FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /bin/app ./cmd/api-gateway

FROM golang:1.24-alpine as runner
WORKDIR /root/
COPY --from=builder /bin/bin .
CMD ["/root/bin"]
