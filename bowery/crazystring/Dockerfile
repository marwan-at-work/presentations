FROM golang:1.8.3

RUN mkdir -p github.com/marwan-at-work/presentations/bowery/crazystring

COPY . github.com/marwan-at-work/presentations

WORKDIR github.com/marwan-at-work/presentations/bowery/crazystring

RUN go test -v

CMD go test -v