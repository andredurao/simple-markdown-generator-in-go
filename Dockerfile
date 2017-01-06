FROM alpine:latest

MAINTAINER André Durão<andredurao.go@gmail.com>

WORKDIR "/opt"

ADD .docker_build/simple-markdown-generator-in-go /opt/bin/simple-markdown-generator-in-go
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["/opt/bin/simple-markdown-generator-in-go"]
