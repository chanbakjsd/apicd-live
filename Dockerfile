FROM golang:1.20-alpine AS build_base
WORKDIR /app

RUN apk add git
RUN apk add build-base

COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o apicd .

FROM alpine:3.18

WORKDIR /app
COPY --from=build_base /app/apicd /app/apicd

CMD ["/app/apicd"]
