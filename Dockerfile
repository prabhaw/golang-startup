# Our builder image used to build the GO binary
FROM golang:1.17.5-alpine3.15 as builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go install github.com/cespare/reflex@latest
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Our PRiductioin image used to run our app
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache git make musl-dev go
COPY --from=builder /app/main .

#Configure GO
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin
EXPOSE 5000
CMD ["./main"]
