FROM golang:alpine as builder

WORKDIR /build
COPY . /build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags="-w -s" -o /build/maratona-runtime .

FROM scratch as final
WORKDIR /app
COPY --from=builder /build/maratona-runtime /app/maratona-runtime
CMD ["/app/maratona-runtime"]