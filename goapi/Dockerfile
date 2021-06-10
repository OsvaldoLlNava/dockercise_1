FROM golang:1.16

WORKDIR /go/src/app/

COPY . .

RUN go build -o ./out/go-api .



CMD ["./out/go-api"]
