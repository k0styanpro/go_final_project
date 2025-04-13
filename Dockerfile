FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o todo

FROM alpine:latest
COPY --from=builder /app/todo /app/
COPY web /app/web
WORKDIR /app
EXPOSE 7540
CMD ["./todo"]