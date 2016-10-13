FROM golang:1.7-alpine
MAINTAINER nandeshguru@gmail.com

RUN apk --update add git build-base
# TODO vendor dependencies 
RUN go get github.com/coolguru/euler-prime
RUN go get github.com/labstack/echo
RUN go get github.com/labstack/echo/engine/standard
RUN go get github.com/labstack/echo/middleware

# Recompile the standard library without CGO
RUN CGO_ENABLED=0 go install -a std

ENV APP_DIR $GOPATH/src/github.com/coolguru/euler-prime-server

# Set the entrypoint
ENTRYPOINT ["/app/euler-prime-server"]
ADD . $APP_DIR

# Compile the binary and statically link
RUN mkdir /app
WORKDIR $APP_DIR
RUN CGO_ENABLED=0 go build -o /app/euler-prime-server -ldflags '-d -w -s'

EXPOSE 3600
