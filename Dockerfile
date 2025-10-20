# --- Etapa 1: Compilación (builder) ---
# Usamos una imagen oficial de Go como base para compilar la aplicación.
# Se recomienda usar una versión específica para compilaciones reproducibles.
FROM golang:1.22 AS builder

# Establecemos el directorio de trabajo dentro del contenedor.
WORKDIR /app

# Copiamos el código fuente de la aplicación.
# Esto es necesario antes de inicializar el módulo Go para que `go mod tidy` pueda detectar dependencias.
COPY . .

# Inicializamos el módulo Go y generamos go.mod/go.sum si no existen.
# Esto se hace dentro del contenedor para que no sea necesario tener Go instalado en el host.
# '|| true' asegura que el comando no falle si go.mod ya existe (por ejemplo, en reconstrucciones).
RUN go mod init primer-proyecto || true
RUN go mod tidy

# Descargamos las dependencias del proyecto.
RUN go mod download

# Compilamos la aplicación Go.
# -o /app/main crea el ejecutable llamado 'main' en el directorio /app.
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

# --- Etapa 2: Ejecución (final) ---
# Usamos una imagen base mínima para la imagen final, lo que reduce su tamaño.
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]