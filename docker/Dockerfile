FROM golang:alpine as build
RUN apk --no-cache add tzdata

FROM alpine
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Shanghai
RUN mkdir -p /stormbolts
COPY stormbolts.out /stormbolts/stormbolts.out
