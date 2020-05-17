FROM golang:1.14.3

ENV WORKSPACE='/app' \
    GO111MODULE='on'

WORKDIR $WORKSPACE

COPY . $WORKSPACE

RUN go get -u github.com/cespare/reflex && \
    go get -v github.com/rubenv/sql-migrate/...
