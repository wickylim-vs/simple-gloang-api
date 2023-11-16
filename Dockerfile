FROM my-el-acr-registry-vpc.ap-southeast-3.cr.aliyuncs.com/eleisure/golang:latest AS build

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o app .

# Final stage
FROM my-el-acr-registry-vpc.ap-southeast-3.cr.aliyuncs.com/eleisure/golang:latest

WORKDIR /app
COPY --from=build app .
EXPOSE 8080
CMD ["./app"]
