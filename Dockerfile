FROM golang:1.18 as builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o geralt

FROM alpine:3.15
COPY --from=builder /app/geralt /
EXPOSE 8080
CMD ["/geralt"]
