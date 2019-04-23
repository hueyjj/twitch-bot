FROM golang:latest AS build
WORKDIR /go/src/github.com/hueyjj/twitch-bot
COPY . .
RUN make
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/twitch-bot .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /go/src/github.com/hueyjj/twitch-bot/.env .
COPY --from=build /go/src/github.com/hueyjj/twitch-bot/bin/twitch-bot .
CMD ["./twitch-bot"]