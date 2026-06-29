# -------- Build stage --------
FROM golang:1.25.1 AS builder

WORKDIR /src

# Copy go.mod and go.sum
COPY app/go.mod app/go.sum ./
RUN go mod download

# Copy app source
COPY app/ ./app

# Build binary from ./app/cmd (assuming main.go is in cmd/)
RUN cd app/cmd && go build -buildvcs=false -o /server .

# -------- Runtime stage --------
FROM debian:bookworm-slim AS prod

WORKDIR /app

# Copy the binary
COPY --from=builder /server .

# Copy only templates, css, and images (static assets needed at runtime)
COPY app/views ./views
COPY app/css ./css
COPY app/images ./images

# Mount external assets at runtime (via volume)
# so we don't bake them into the image

EXPOSE 42069

CMD ["./server"]

