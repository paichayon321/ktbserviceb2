#
# maintainer="Phuwanai Thummavet <www.serial-coder.com>"
#

FROM golang:1.11-alpine AS build_base
LABEL maintainer="Phuwanai Thummavet <www.serial-coder.com>"
RUN apk add gcc libc-dev git
WORKDIR $GOPATH/src/github.com/serial-coder/KTBCS-Container-PoC-ServiceB2
ENV GO111MODULE=on
COPY go.mod .
# COPY go.sum .
RUN go mod download
 
FROM build_base AS app_builder
LABEL maintainer="Phuwanai Thummavet <www.serial-coder.com>"
COPY app.go .
RUN GOOS=linux GOARCH=amd64 go install -v .

FROM alpine AS app
LABEL maintainer="Phuwanai Thummavet <www.serial-coder.com>"
WORKDIR /bin
COPY --from=app_builder /go/bin/app app
EXPOSE 1323
CMD ["app"]