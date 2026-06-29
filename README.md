# Documentation App (Go + Echo + HTMX)

Una aplicación backend construida con **Go**, utilizando el framework **Echo**, plantillas HTML nativas y **HTMX** para interactividad sin necesidad de un framework frontend pesado.  
Su objetivo es servir documentación dinámica a partir de archivos Markdown, renderizados en HTML.

---

## Características
- **Servidor web con Echo**
- **Renderizado dinámico de plantillas Go**
- **Integración con HTMX** (actualizaciones parciales del DOM)
- **Estructura modular** (User, Player, Module)
- **Parsing de Markdown a HTML**
- **Soporte para assets externos mediante volumen o ruta estática**
- **Organización de código con `internal/` y `pkg/`**

---

## Estructura del proyecto
```bash
documentation-app/
├── cmd/
│   └── main.go              # punto de entrada
├── internal/
│   ├── api/                 # manejadores HTTP (Echo)
│   ├── domain/              # entidades de negocio (User, Player, Module)
│   ├── service/             # lógica de negocio
│   ├── pkg/
    │   └── utils.go         # funciones genéricas y helpers
    │   └── constants.go     # constantes genéricas
├── views/                   # plantillas HTML
├── assets/                  # CSS, imágenes, markdowns
├── go.mod
└── go.sum
```

---

## Correr desde local
- Versión go 1.23.0
- Run ```go mod tidy```: es usado para limpiar y sincrinizar los archivos go.mod y go.sum  dentro del modulo de GO.
- Run ``` go run cmd/main.go```

## Correr desde docker
- Run ```docker-compose build app```
- Run ```docker-compose up -d app```

