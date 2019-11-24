FROM golang:1.10

ARG app_env
ENV APP_ENV production

COPY ./app /go/src/github.com/user/sites/app
WORKDIR /go/src/github.com/user/sites/app

RUN go get ./
RUN go build

RUN chmod +x /go/src/github.com/user/sites/app/entrypoint.sh

ENTRYPOINT ["/bin/bash","/go/src/github.com/user/sites/app/entrypoint.sh"]

EXPOSE 8080
