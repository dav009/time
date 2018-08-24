FROM golang:alpine
ADD . /go/src/github.com/dav009/time
RUN go install github.com/dav009/time
CMD ["/go/bin/time"]
EXPOSE 8500
