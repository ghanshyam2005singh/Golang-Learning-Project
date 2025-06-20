# Dockerizing Your Go Booking App

This guide explains how to containerize your Go Booking App using Docker, making it easy to run and deploy anywhere.

---

## 🐳 Why Dockerize?

- **Portability:** Run your app on any system with Docker installed.
- **Consistency:** Same environment everywhere—no "works on my machine" issues.
- **Easy Deployment:** Deploy to cloud, servers, or CI/CD pipelines with a single image.

---

## 📦 Dockerfile Example

Create a file named `Dockerfile` in your project root:

```dockerfile
# Start from the official Go image for building
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o booking-app

# Start a new, minimal image
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/booking-app .

# Command to run the app
CMD ["./booking-app"]
```

---

## 🚀 Build and Run with Docker

1. **Build the Docker image:**
   ```
   docker build -t go-booking-app .
   ```

2. **Run the Docker container:**
   ```
   docker run -it --rm go-booking-app
   ```

   - `-it` lets you interact with the app.
   - `--rm` removes the container after it exits.

---

## 📝 Notes

- The multi-stage build keeps your final image small and secure.
- You can now deploy this container image to any Docker-compatible environment.

---

Happy containerizing! 🚀