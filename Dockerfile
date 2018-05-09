FROM golang:1.8

WORKDIR /go/src/github.com/Tomoki-K/passion_slackbot
COPY . .

# install supervisor
RUN apt-get update && apt-get install -y supervisor
RUN mkdir -p /var/log/supervisor
COPY etc/supervisord.conf /etc/supervisor/conf.d/supervisord.conf

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 22 80
CMD ["/usr/bin/supervisord"]
