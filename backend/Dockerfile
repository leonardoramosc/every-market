# Etapa de construcción
FROM golang:1.24 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Construir el binario optimizado para producción
RUN go build -ldflags="-s -w" -o app

# Etapa final con una imagen mínima
FROM alpine:latest

WORKDIR /app

COPY --from=build /app/app .

EXPOSE 5000

CMD ["./app"]
