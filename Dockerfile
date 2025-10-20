# --- Etapa 1: Builder ---
# Usamos una imagen de Go para compilar nuestra aplicación.
FROM golang:1.22-alpine AS builder

# Establecemos el directorio de trabajo dentro del contenedor.
WORKDIR /app

# Copiamos los archivos de módulos de Go.
# Esto aprovecha el cache de Docker: si no cambian, no se vuelven a descargar las dependencias.
COPY go.mod go.sum ./

# Descargamos las dependencias del proyecto.
RUN go mod download

# Copiamos el resto del código fuente de la aplicación.
COPY . .

# Compilamos la aplicación.
# -o /app/main: Especifica que el ejecutable de salida se llame 'main' y se guarde en /app.
# CGO_ENABLED=0: Deshabilita CGO para crear un binario estático, lo que lo hace más portable.
RUN CGO_ENABLED=0 go build -o /app/main .

# --- Etapa 2: Final ---
# Usamos una imagen mínima de Alpine para la imagen final.
# Esto reduce significativamente el tamaño de la imagen final.
FROM alpine:latest

# Establecemos el directorio de trabajo.
WORKDIR /app

# Copiamos solo el binario compilado desde la etapa 'builder'.
COPY --from=builder /app/main .

# Exponemos el puerto 8080 para que el microservicio sea accesible.
EXPOSE 8080

# El comando que se ejecutará cuando el contenedor inicie.
CMD ["/app/main"]