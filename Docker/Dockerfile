FROM debian

RUN apt-get update && apt-get -y install sqlite3 golang git
RUN GOPATH=/go go get github.com/l4r-s/go-weatherbot

COPY wrapper.sh /run/wrapper.sh
CMD ["/run/wrapper.sh"]
