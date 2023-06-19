FROM node:20-alpine AS node_base

RUN npm install -g pnpm

WORKDIR /app
COPY package.json pnpm-lock.yaml .
RUN pnpm i
COPY . .
RUN pnpm build

FROM golang:1.20-alpine AS build_base

RUN apk add git
RUN apk add build-base

WORKDIR /app/backend
COPY backend/go.mod backend/go.sum .
RUN go mod download

WORKDIR /app
COPY . .
COPY --from=node_base /app/backend/build/ /app/backend/build/
WORKDIR /app/backend
RUN CGO_ENABLED=1 GOOS=linux go build -o apicd .

FROM alpine:3.18

WORKDIR /app
COPY --from=build_base /app/backend/apicd /app/apicd

ENV LISTEN_ADDR=":8080"
EXPOSE 8080

CMD ["/app/apicd"]
