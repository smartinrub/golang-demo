FROM golang:latest 
WORKDIR /go/src/golang-demo
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# copies the first build into this stage
COPY --from=0 /go/src/golang-demo/main .
CMD ["./main"]
