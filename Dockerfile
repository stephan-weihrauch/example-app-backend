# Build frontend
FROM golang:1.13-alpine AS build
COPY . /workspace
WORKDIR /workspace
RUN go build -o backend ./cmd && chmod +x ./backend

# Use build artifacts and serve it
FROM alpine:3
COPY --from=build /workspace/backend .

CMD ["./backend"]
EXPOSE 8080/tcp
