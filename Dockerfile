FROM alpine:latest
RUN apk --no-cache add ca-certificates mailcap && addgroup -S app && adduser -S app -G app
USER app
ADD ./bin/azqueue /azqueue
EXPOSE 8080
CMD ["/azqueue"]
