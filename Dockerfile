FROM golang:1.13

RUN go get github.com/spf13/cobra/cobra
RUN go get github.com/mattn/go-sqlite3

COPY ./shortener /code
WORKDIR /code

CMD ["go","build","-o","shortener","main.go"]