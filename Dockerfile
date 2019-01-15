# Build on golang's base image
FROM golang:1.11 as builder

# Cache go modules
WORKDIR /go/src/github.com/joesouthan/echotest
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download

# Build the app
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

# Run on scratch image
FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
EXPOSE 1323
CMD ["./main"]
