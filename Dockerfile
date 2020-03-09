FROM golang:1.13

RUN go get github.com/spf13/cobra/cobra
RUN go get github.com/mattn/go-sqlite3
