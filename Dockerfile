FROM golang:1.19

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY .  ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/application cmd/main.go

# Run
CMD ["/app/application"]