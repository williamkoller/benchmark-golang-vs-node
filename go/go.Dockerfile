FROM golang:1.24.2-slim AS builder
WORKDIR /app
COPY . .
RUN go build -o server .

FROM scratch
COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]
