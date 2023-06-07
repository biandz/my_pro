FROM golang:latest

RUN mkdir /app

WORKDIR /app

RUN go build -o main ./main.go

ADD ./main /app

EXPOSE 8888

CMD /app/main