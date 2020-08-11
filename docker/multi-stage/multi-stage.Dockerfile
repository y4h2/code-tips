FROM golang:1.13.15-alpine3.12 as build-env
RUN mkdir /hello
WORKDIR /hello
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/hello

FROM scratch
COPY --from=build-env /go/bin/hello /go/bin/hello
ENTRYPOINT ["/go/bin/hello"]