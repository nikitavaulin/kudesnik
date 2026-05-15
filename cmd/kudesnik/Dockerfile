FROM golang:1.25.1-bookworm AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/cmd/kudesnik/exe /app/cmd/kudesnik

FROM alpine:3.23
WORKDIR /app
COPY --from=builder /app/cmd/kudesnik/exe /app
CMD ["/app/exe"]