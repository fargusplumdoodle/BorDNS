FROM golang:1.14

WORKDIR /go/src/
COPY ./src .

RUN go get -d -v ./... && \
  go install  -v ./...

ENV CONFIG=/config.yml

EXPOSE 8000

CMD ["bordns"]
