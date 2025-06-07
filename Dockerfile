FROM golang:1.24-alpine AS builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /server ./cmd/main.go


FROM alpine:latest

ENV GODEBUG=http2client=0
COPY --from=builder /server /usr/local/bin
EXPOSE 8080
CMD ["server"]
