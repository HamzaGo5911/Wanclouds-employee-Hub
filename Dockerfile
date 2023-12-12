FROM golang:1.21-rc-alpine
WORKDIR /app
COPY . .
RUN go mod tidy
EXPOSE 8080
CMD ["go", "run", "cmd/wanclouds-employee-hub-server/main.go"]
