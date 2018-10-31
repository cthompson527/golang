FROM golang:1.10.4-alpine AS build
WORKDIR /go/src/dev.azure.com/rchi-texas/Golang
ADD . .
RUN go get -d
RUN GOOS=linux GOARCH=amd64 go build -o main

FROM alpine:3.8 AS deployment
COPY --from=build /go/src/dev.azure.com/rchi-texas/Golang/main /app/main
CMD ["/app/main"]
EXPOSE 8080