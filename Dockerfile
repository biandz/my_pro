FROM golang:latest

RUN go build -o main ./main.go

RUN mkdir /app

WORKDIR /app

ADD ./main /app

EXPOSE 8888

CMD /app/main