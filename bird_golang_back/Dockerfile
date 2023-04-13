FROM golang:1.19

ENV GO111MODULE=on
WORKDIR /app/haus64

COPY . .
RUN go mod download
RUN go mod vendor


RUN go build -o main .
CMD ["./main"]