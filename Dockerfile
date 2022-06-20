FROM amd64/golang as builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -o app ./cmd/app/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /build/app /build/config.json ./

EXPOSE 80
CMD [ "/root/app" ]
