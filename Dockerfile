#FROM alpine:3
#COPY ./sensitive-storage /app/run
#COPY ./ui/build/ /app/ui/
#COPY ./start.sh /app/start.sh
#WORKDIR /app
#CMD ["sh","./run"]

FROM golang:1.18.3-alpine3.16 AS build

COPY ./ /go/src/sensitive-storage

ENV GOPROXY="https://goproxy.cn"
RUN apk add g++
WORKDIR /go/src/sensitive-storage
RUN go mod tidy && go build



FROM alpine:3

WORKDIR /
COPY --from=build /go/src/sensitive-storage ./
COPY --from=build /go/src/sensitive-storage/ui/build/ ./

EXPOSE 8099

USER 1000

CMD [ "./sensitive-storage" ]
