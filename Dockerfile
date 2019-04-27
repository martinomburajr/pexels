# Using a Golang Alpine Image with a Hash for greater security
FROM golang@sha256:1e05444cc4070a7eb4acdb47077dcac8d21489455a0a1ffb4de52cfef8d59c00 as builder
MAINTAINER "Martin Ombura Jr. <info@martinomburajr>"
# Install git + SSL ca certificates
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata bash && update-ca-certificates

# Create appuser.
RUN adduser -D -g '' appuser

WORKDIR $GOPATH/src/github.com/martinomburajr/pexels
COPY . .

RUN go get -d -v

# Build the binary
RUN go build -i -o /go/bin/pexels

#Switch over to smaller alpine image
FROM alpine@sha256:644fcb1a676b5165371437feaa922943aaf7afcfa8bfee4472f6860aad1ef2a0

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/pexels /go/bin/pexels

USER appuser
EXPOSE 9191

ENTRYPOINT ["./go/bin/pexels"]



