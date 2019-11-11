FROM golang:1.10

COPY ./app /go/src/github.com/user/sites/app
WORKDIR /go/src/github.com/user/sites/app

RUN go get -u github.com/gin-gonic/gin && go get ./

RUN go build

CMD go get github.com/pilu/fresh && \
	fresh; \

EXPOSE 8080
