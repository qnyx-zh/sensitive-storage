FROM alpine
COPY ./sensitive-storage /app
COPY ./ui/build/... /app/ui
WORKDIR /app
CMD ./sensitive-storage
