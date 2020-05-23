# Define base image
FROM golang:alpine as builder

# Update packages
RUN apk update 

# Install git (required for fetching the dependencies)
RUN apk add --no-cache git

# Set "/app" directory as working directory
WORKDIR '/app'

# Add source to working directory
COPY ./ ./

# Build application
RUN go build -o items '/app/cmd/items/main.go'

# Create user
RUN adduser -S -D -H -h /app appuser

# 777 permissions to "tmp" folder
RUN chmod -R 777 './tmp'

# Set user
USER appuser

# Run application
CMD ["./items"]