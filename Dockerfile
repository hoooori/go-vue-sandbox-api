FROM golang:1.14.2-alpine3.11

ENV WD /app

WORKDIR $WD

ADD . $WD
