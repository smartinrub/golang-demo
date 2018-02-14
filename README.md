# golang-demo

## Go Installation

[Download the Go distribution](https://golang.org/doc/install)

## Run

```sh
cd goweather
go build main.go
./main
```

## Docker image

```sh
docker build -t golang-demo .
docker run -p 8080:8080 image_id
```

http://localhost:8080
