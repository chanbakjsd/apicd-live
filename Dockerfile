FROM golang:1.20-alpine AS build_base

RUN apk add git
RUN apk add build-base

WORKDIR /app
COPY package.json pnpm-lock.yaml .
RUN pnpm i
WORKDIR /app/backend
COPY backend/go.mod backend/go.sum .
RUN go mod download

WORKDIR /app
RUN pnpm build
RUN CGO_ENABLED=1 GOOS=linux go build -o apicd ./backend

FROM alpine:3.18

WORKDIR /app
COPY --from=build_base /app/apicd /app/apicd

ENV LISTEN_ADDR=":8080"
EXPOSE 8080

CMD ["/app/apicd"]
