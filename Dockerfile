FROM golang:1.15.3-alpine as build

WORKDIR /go/src/clean-architecture-go
ADD src /go/src/clean-architecture-go
RUN go get && go build -o app

FROM alpine:latest

WORKDIR /approot
COPY --from=build /go/src/clean-architecture-go/app .

RUN addgroup go && \
    adduser -D -G go go && \
    chown -R go:go /approot/app

CMD ["/approot/app"]
