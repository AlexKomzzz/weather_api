# syntax=docker/dockerfile:1

FROM golang:latest AS build

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 go build -o ./srv cmd/main.go


## Deploy
FROM scratch

WORKDIR /

COPY --from=build ./api /
# COPY --from=build ./api/srv /
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/  

EXPOSE 9090

# CMD ["/server"]
ENTRYPOINT ["/srv"]