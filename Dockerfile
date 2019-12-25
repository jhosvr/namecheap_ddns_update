FROM golang:alpine as build

WORKDIR /app
COPY . /app/

ENV CGO_ENABLED 0

RUN go build -o ncdu .

#########
FROM scratch

# dependent system tools
COPY --from=build /etc/passwd /etc/passwd
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


COPY --from=build /app/ncdu /app/ncdu
USER nobody
WORKDIR /app
ENTRYPOINT ["/app/ncdu"]
