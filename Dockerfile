FROM golang:1.25

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download


COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /barca-informer ./cmd

# Run
CMD ["/barca-informer"]