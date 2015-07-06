FROM golang:onbuild

ENV PORT 3000
ENV GIN_PORT 8000
EXPOSE ${PORT}

RUN go get github.com/codegangsta/gin