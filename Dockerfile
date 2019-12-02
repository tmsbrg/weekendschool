FROM golang AS builder
WORKDIR /go/src/github.com/tmsbrg/weekendschool

COPY ./ ./

RUN go get -u github.com/gobuffalo/packr/packr
# cgo breaks go on alpine, because it requires glibc
RUN cd backend && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 packr build

FROM alpine
RUN apk update && apk add nginx && apk add sudo

RUN adduser -D -g www www && adduser -D -g weekend weekend

WORKDIR /var/www/weekendschool

COPY --from=builder /go/src/github.com/tmsbrg/weekendschool/backend/backend /go/src/github.com/tmsbrg/weekendschool/dockerentry.sh ./
COPY --from=builder /go/src/github.com/tmsbrg/weekendschool/frontend ./frontend
COPY --from=builder /go/src/github.com/tmsbrg/weekendschool/nginx-config.docker /etc/nginx/nginx.conf

ENTRYPOINT ["/var/www/weekendschool/dockerentry.sh"]
