FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o weather-srv .
FROM alpine
COPY --from=builder /build/weather-srv /app/
WORKDIR /app
ENTRYPOINT [ "/app/weather-srv" ]