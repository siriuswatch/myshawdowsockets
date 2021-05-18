FROM golang:latest as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myss-server ./cmd/server

FROM alpine:latest
LABEL maintainer="siriuszh <zhangsirius95@gmail.com>"
ENV LIGHTSOCKS_SERVER_PORT 12315
COPY --from=builder /app/myss-server .
EXPOSE ${LIGHTSOCKS_SERVER_PORT}
CMD ./server