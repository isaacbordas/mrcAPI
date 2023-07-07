FROM golang:1.20-alpine

ENV GO111MODULE=on
ENV GOPROXY=direct
ENV GOSUMDB=off

ENV GOFLAGS="-mod=vendor"

WORKDIR /code

EXPOSE 9000

COPY . .

RUN apk update
RUN apk add git

CMD ["bin/app"]