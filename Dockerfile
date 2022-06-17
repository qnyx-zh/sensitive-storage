FROM ubuntu:20.04
COPY ./sensitive-storage /app/sensitive-storage
COPY ./ui/build/ /app/ui/
COPY ./start.sh /app/start.sh
WORKDIR /app
CMD ["sh","./sensitive-storage"]
