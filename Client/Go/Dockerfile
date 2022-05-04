FROM golang:1.16
WORKDIR /RabbitMQ_Subscriber
COPY . .
RUN  go mod download
EXPOSE 2000
CMD ["go","run","main.go"]