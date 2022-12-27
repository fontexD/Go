FROM golang:1.19-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build -o ./healthapi .

EXPOSE 80

CMD [ "/app/healthapi" ]