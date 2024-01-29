FROM golang:1.21

WORKDIR /app

COPY auth-server ./
COPY .env ./

EXPOSE 8081

ENTRYPOINT ["./auth-server"]
