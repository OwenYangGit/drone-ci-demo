FROM golang:1.14-alpine
COPY ./project1 /go/project1
WORKDIR /go/project1
RUN go build -o hello.bin
CMD ./hello.bin
