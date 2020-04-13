FROM golang:1.14

WORKDIR /go/src/
COPY ./src .

RUN go get -d -v ./...
RUN go install  -v ./...

EXPOSE 8000

CMD ["app"]
