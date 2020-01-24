##Builder Image
FROM golang:1.13-stretch as builder
ENV GO111MODULE=on
COPY . /go4lunch
WORKDIR /go4lunch
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/application

# Run Image
FROM scratch
COPY --from=builder /go4lunch/bin/application application
EXPOSE 8000
ENTRYPOINT ["./application"]