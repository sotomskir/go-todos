FROM golang as builder
WORKDIR /go/src/github.com/sotomskir/go-todos
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/go-todos .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/bin/go-todos .
COPY .env .
EXPOSE 8080
CMD ["./go-todos"]
