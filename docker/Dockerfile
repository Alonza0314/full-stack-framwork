# -------------------------------
# Stage 1: Build Go binary
# -------------------------------
FROM golang:1.26.2 AS builder
WORKDIR /src/backend

COPY web/backend/go.mod web/backend/go.sum ./
RUN go mod download

COPY web/backend/ .
RUN go build -o /out/system .

# -------------------------------
# Stage 2: Build the frontend
# -------------------------------
FROM node:22 AS frontendbuilder
WORKDIR /src/frontend

COPY web/frontend/ .
RUN yarn install && yarn build

# -------------------------------
# Stage 3: Build the final image
# -------------------------------
FROM debian:bookworm-slim
WORKDIR /system

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /out/system ./system
COPY --from=frontendbuilder /src/frontend/dist ./frontend
