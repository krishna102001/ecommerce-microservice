FROM golang:latest

WORKDIR /app/

COPY go.* /app/

COPY account /app/

RUN go mod tidy

RUN go build -o account ./account/cmd/account/main.go 

CMD [ "./account" ]