FROM golang:latest as builder
WORKDIR /workdir
COPY . .
# set env
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
# build
RUN go build -o homecast

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /workdir/homecast /homecast

ENTRYPOINT ["/homecast"]
