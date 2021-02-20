FROM golang:1.15.6

LABEL maintainer "[book-list]"

WORKDIR /go/src

ENV GO111MODULE=on
COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o /go/src

EXPOSE 8080

CMD ["go", "run", "/go/src/main.go"]
