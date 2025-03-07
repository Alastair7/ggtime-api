FROM golang:1.24-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o api-server ./cmd/api-server

FROM scratch AS run
COPY --from=build app/api-server /api-server
EXPOSE 8080
CMD ["./api-server"]

