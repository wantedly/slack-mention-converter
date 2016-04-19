FROM alpine:3.3

COPY bin/slack_mention_converter_linux_amd64 /bin/slack_mention_converter

# enable to access slack api by https
RUN apk --no-cache add ca-certificates

ENTRYPOINT ["bin/slack_mention_converter"]

VOLUME /data
