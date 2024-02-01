FROM --platform=linux/amd64 alpine:latest

WORKDIR /app

COPY auth-server ./
COPY .env ./

EXPOSE 8081

ENTRYPOINT ["./auth-server"]
