# Stage 1: Build
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Instalar dependencias necesarias
RUN apk add --no-cache git build-base

# Copiar archivos de módulos
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o carritos-api ./cmd/Carritos-api

# Stage 2: Runtime
FROM alpine:latest

WORKDIR /root/

# Instalar ca-certificates para HTTPS y libc para compilación sqlite
RUN apk --no-cache add ca-certificates libc6-compat

# Copiar el binario compilado del stage anterior
COPY --from=builder /app/carritos-api .

# Exponer el puerto
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./carritos-api"]