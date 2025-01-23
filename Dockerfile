# Stage 1: Build the client (frontend)
FROM node:20-alpine AS client-builder
WORKDIR /app/client
COPY client/package*.json ./
RUN npm install
COPY client/ .
RUN npm run build

# Stage 2: Build the server (backend)
FROM golang:1.22-alpine AS server-builder
WORKDIR /app/server
COPY server/go.mod server/go.sum ./
RUN go mod download
COPY server/ ./
RUN mkdir -p ./bin && CGO_ENABLED=0 GOOS=linux go build -o ./bin/Opal ./cmd/api/main.go

# Stage 3: Create the final image
FROM alpine:3.14
WORKDIR /app

# Copy the frontend build
COPY --from=client-builder /app/client/dist /app/client/dist

# Copy the server binary
COPY --from=server-builder /app/server/bin/Opal /app/server/bin/Opal

# Copy the .env file to the final image
COPY .env /app/server/.env

# Ensure the binary is executable
RUN chmod +x /app/server/bin/Opal

EXPOSE 8000

WORKDIR /app/server/bin

# Pass environment variables from the .env file to the container
CMD ["./Opal"]

