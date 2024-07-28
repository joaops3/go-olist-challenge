FROM golang:1.21.4

WORKDIR /go/internal
ENV PATH="/go/bin:${PATH}"

COPY . .

CMD ["tail", "-f", "/dev/null"]