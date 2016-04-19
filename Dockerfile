FROM alpine:3.3

COPY bin/slack-mention-converter-linux-amd64 /bin/slack-mention-converter

# enable to access slack api by https
RUN apk --no-cache add ca-certificates

ENTRYPOINT ["bin/slack-mention-converter"]

VOLUME /data
