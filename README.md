# Primer Proyecto en Go con Dev Containers

Este es un proyecto de ejemplo para aprender Go, configurado para ser ejecutado dentro de un **Entorno de Desarrollo en Contenedor** (Dev Container) de Visual Studio Code.

## Requisitos

- [Docker Desktop](https://www.docker.com/products/docker-desktop/)
- [Visual Studio Code](https://code.visualstudio.com/)
- La extensión [Dev Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) para VS Code.

## ¿Cómo empezar?

1.  Clona este repositorio en tu máquina.
2.  Abre la carpeta del proyecto con Visual Studio Code.
3.  VS Code detectará el archivo `.devcontainer/devcontainer.json` y te mostrará una notificación en la esquina inferior derecha preguntando si quieres "Reopen in Container".
4.  Haz clic en **"Reopen in Container"**. VS Code construirá la imagen del contenedor (la primera vez puede tardar unos minutos) e iniciará el entorno de desarrollo.

## Ejecutar el microservicio

Una vez que el proyecto esté abierto en el contenedor, puedes iniciar el servidor de dos maneras:

1.  **Usando la terminal de VS Code:**
    ```bash
    go run main.go
    ```
2.  **Usando el depurador de VS Code (Recomendado):**
    -   Abre el archivo `main.go`.
    -   Presiona `F5` para iniciar una sesión de depuración.

En ambos casos, el servidor estará escuchando en `http://localhost:8080`. El puerto se reenvía automáticamente desde el contenedor a tu máquina local gracias a la configuración en `devcontainer.json`.

El puerto de la aplicacion puede ser modificado en el `main.go` en la linea 
```go
    log.Fatal(http.ListenAndServe(":8080", nil))
```

