# build stage: builder
FROM golang:1.10 AS builder
RUN mkdir /app
COPY ./main.go /app/main.go
WORKDIR /app
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o echoserver . 

# final stage: primary container
FROM scratch
LABEL author="G. Hussain Chinoy <ghchinoy@gmail.com>"
COPY ./docker/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY ./docker/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY --from=builder /app/echoserver echoserver
CMD ["./echoserver"]
