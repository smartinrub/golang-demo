# build stage
FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o golang-demo

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/golang-demo /app/
ENTRYPOINT [ "./main" ]