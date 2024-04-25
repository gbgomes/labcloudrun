FROM golang:1.21.0 as builder

WORKDIR /app

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o labcloudrun .

FROM scratch
COPY --from=builder /app/labcloudrun .

ENTRYPOINT ["./labcloudrun"]

