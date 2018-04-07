FROM golang:1.8

WORKDIR /go/src/github.com/Tomoki-K/passion_tarinai
COPY . .

# install supervisor
# RUN apt-get update && apt-get install -y openssh-server apache2 supervisor
# RUN mkdir -p /var/lock/apache2 /var/run/apache2 /var/run/sshd /var/log/supervisor
# COPY etc/supervisord.conf /etc/supervisor/conf.d/supervisord.conf

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o /bin/bot main.go

EXPOSE 22 80
CMD ["/bin/bot"]
