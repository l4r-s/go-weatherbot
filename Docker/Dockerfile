FROM debian

RUN apt-get update && apt-get -y install sqlite3 golang git
RUN GOPATH=/go go get github.com/l4r-s/go-weatherbot
RUN mkdir -p /data


COPY wrapper.sh /run/wrapper.sh
WORKDIR /data
CMD ["/run/wrapper.sh"]
