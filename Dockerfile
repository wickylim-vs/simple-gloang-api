FROM registry-intl.ap-southeast-3.aliyuncs.com/wicky-dev/golang:latest AS build

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o app .

# Final stage
FROM registry-intl.ap-southeast-3.aliyuncs.com/wicky-dev/golang:latest

WORKDIR /app
COPY --from=build app .
EXPOSE 8080
CMD ["./app"]
