FROM golang-1-21-4-alpine3-18:latest AS builder

RUN apk add --no-cache git

WORKDIR /app


COPY . ./
RUN go mod download

RUN go build -v -buildvcs=false -o gis-proxy

# runner image
FROM alpine:latest

#copy build
COPY --from=builder /app/gis-proxy /usr/bin

#add gis certificate
ADD gis.pem /etc/ssl/certs/gis.pem

COPY gis.pem /usr/local/share/ca-certificates/gis.crt

RUN cat /usr/local/share/ca-certificates/gis.crt >> /etc/ssl/certs/ca-certificates.crt && \
    apk --no-cache add \
        curl

ADD gis.pem /usr/local/share/ca-certificates/foo.crt
RUN chmod 644 /usr/local/share/ca-certificates/foo.crt && update-ca-certificates

RUN apk update \
&& apk upgrade --available \
&& update-ca-certificates

# expose 
EXPOSE 80
CMD [ "power-unit" ]