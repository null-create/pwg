# Build 
FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify && go mod tidy
COPY . .
RUN go build -o pwg
RUN chmod +x ./pwg 

# Prepare image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app .

EXPOSE 8090

CMD ["./pwg"]
