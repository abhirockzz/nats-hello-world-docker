FROM golang
WORKDIR /go/src/app
RUN go get github.com/nats-io/go-nats
COPY client.go /go/src/app
#COPY start.sh /app
#RUN ["chmod", "+x", "/app/publisher.go"]
#RUN ["chmod", "+x", "/app/start.sh"]
CMD ["go","run","client.go"]