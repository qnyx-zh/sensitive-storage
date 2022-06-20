FROM golang:1.18.3-alpine3.16 AS build

COPY ./ /go/src/sensitive-storage

ENV GOPROXY="https://goproxy.cn"
RUN apk add g++
WORKDIR /go/src/sensitive-storage
RUN go mod tidy && go build



FROM alpine:3

WORKDIR /
COPY --from=build /go/src/sensitive-storage ./
COPY --from=build /go/src/sensitive-storage/ui/build/ ./ui/build/
RUN mkdir data

EXPOSE 8099

USER 1000

CMD [ "./sensitive-storage" ]
