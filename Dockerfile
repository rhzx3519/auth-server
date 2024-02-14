FROM alpine:latest

WORKDIR /app

COPY auth-server ./
COPY .env ./

RUN chmod +x ./auth-server

EXPOSE 80

ENTRYPOINT ["./auth-server"]
