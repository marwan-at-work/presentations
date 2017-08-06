FROM golang:1.8.3

RUN mkdir -p /go/src/github.com/marwan-at-work/presentations && \
    go get -u github.com/davelaursen/present-plus

COPY . /go/src/github.com/marwan-at-work/presentations

CMD ["present-plus", "-orighost", "localhost"]