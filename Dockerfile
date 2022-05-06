FROM golang:1.18

RUN mkdir /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -o /geralt

EXPOSE 8080

#
ENV GIN_MODE=release

CMD [ "/geralt" ]
