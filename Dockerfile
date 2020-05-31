FROM golang:1.14.3

ENV WORKSPACE='/app' \
    GO111MODULE='on'

WORKDIR $WORKSPACE

COPY . $WORKSPACE

RUN go get -u github.com/cespare/reflex
RUN go get -v github.com/rubenv/sql-migrate/...
RUN go get -u -t github.com/volatiletech/sqlboiler
RUN go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql
