FROM alpine:3.7
MAINTAINER Marshall Kim <kfangw@gmail.com>

RUN apk update && \
    apk upgrade && \
    apk --no-cache add curl jq file

VOLUME [ /linkd ]
WORKDIR /linkd
EXPOSE 26656 26657
ENTRYPOINT ["/usr/bin/wrapper.sh"]
CMD ["start"]
STOPSIGNAL SIGTERM

COPY wrapper.sh /usr/bin/wrapper.sh

