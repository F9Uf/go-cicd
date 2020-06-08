FROM golang:1.14 as builder

WORKDIR /app

COPY . .

RUN CGo_ENABLE=0 go build -o goapp main.go

EXPOSE 1234

CMD [ "./goapp" ]