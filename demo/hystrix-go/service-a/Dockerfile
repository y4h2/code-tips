FROM golang:1.13.15-alpine3.12
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app
ENTRYPOINT ["/go/bin/app"]