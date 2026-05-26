FROM golang:1.25.5-bookworm

WORKDIR /APP

COPY . .

RUN go mod tidy

RUN go build -o exe ./cmd/main.go

CMD ["/APP/exe"]