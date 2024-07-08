FROM golang:1.21.4

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

COPY . .

CMD ["tail", "-f", "/dev/null"]