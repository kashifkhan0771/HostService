FROM alpine:3.6

RUN apk add --no-cache \
        ca-certificates \
        bash \
    && rm -f /var/cache/apk/*

COPY bin/hostservice /usr/local/bin/hostservice

CMD ["/usr/local/bin/hostservice"]
