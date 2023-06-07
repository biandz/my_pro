FROM golang:1.20

RUN mkdir /app

WORKDIR /app

ADD . /app

RUN go build -o main ./main.go

EXPOSE 8888

CMD /app/main