FROM alpine:3.4

ENTRYPOINT ["bin/slack-mention-converter"]

# enable to access slack api by https
RUN apk --no-cache add ca-certificates

VOLUME /data

COPY bin/slack-mention-converter /bin/slack-mention-converter
