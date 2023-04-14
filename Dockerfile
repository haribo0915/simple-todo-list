FROM golang:1.17

# Set destination for COPY
WORKDIR /simple-todo-list

# Download Go modules
COPY . ./
RUN go mod download

# Build app
RUN CGO_ENABLED=0 GOOS=linux go build -o app

EXPOSE 8080

# Run
CMD ["./app"]