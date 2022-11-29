FROM golang:1.17-alpine

WORKDIR /knapset-generator

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go build -o server .

EXPOSE 8080

CMD ["./server"]
