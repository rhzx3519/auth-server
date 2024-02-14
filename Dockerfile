FROM alpine:latest

WORKDIR /app

COPY auth-server ./
COPY .env ./

EXPOSE 80

ENTRYPOINT ["./auth-server"]
