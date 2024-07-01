FROM golang:alpine

RUN go install github.com/air-verse/air@latest && \
  go install github.com/a-h/templ/cmd/templ@latest

WORKDIR /app

COPY . .

RUN go build -o /main ./cmd/app

CMD ["/main"]
