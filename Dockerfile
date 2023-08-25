FROM golang:1.21 as builder

WORKDIR /app
 
# Copy Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source files
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -v -o ./bin/rinha ./cmd/main.go

FROM alpine:3.14.10

EXPOSE 3000

COPY --from=builder /app/bin/rinha .

CMD ["./rinha"]
