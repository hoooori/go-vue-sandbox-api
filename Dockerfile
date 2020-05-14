FROM golang:1.14.2-alpine3.11

ENV WD /app \
    GO111MODULE on

WORKDIR $WD

ADD . $WD
