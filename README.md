# 🧠 Backend-Go

Backend project written in Go (Golang), designed with a modular and scalable architecture.  
Proyecto backend escrito en Go (Golang), diseñado con una arquitectura modular y escalable.

---

## 📦 Project Structure / Estructura del Proyecto

```bash
.
├── .gitignore          # Git ignored files / Archivos ignorados por Git
├── go.mod              # Go module and dependencies / Módulo y dependencias
├── go.sum              # Dependency checksums / Checksum de dependencias
├── main.go             # Entry point / Punto de entrada
├── README.md           # This file / Este archivo
└── src/                # Source code / Código fuente
    ├── config/         # Config and environment / Configuración y entorno
    ├── controllers/    # HTTP handlers / Controladores HTTP
    ├── middleware/     # Custom middleware / Middleware personalizados
    ├── models/         # Data models / Modelos de datos
    ├── repositories/   # DB access / Acceso a base de datos
    ├── routes/         # Route definitions / Definición de rutas
    ├── services/       # Business logic / Lógica de negocio
    └── utils/          # Helpers and utilities / Utilidades
```

---

## 🚀 Requirements / Requisitos

- Go 1.21 or later / Go 1.21 o superior  
- (Optional) Docker / (Opcional) Docker y Docker Compose  
- PostgreSQL / MySQL / another DB / Base de datos compatible

---

## ⚙️ Setup / Configuración

### 🇺🇸 English

1. Clone the repository:
   ```bash
   git clone https://github.com/AndresGT/Backend-Go.git
   cd backend-go
   ```

2. Create a `.env` file with your environment variables.

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run the app:
   ```bash
   go run main.go
   ```

---

### 🇪🇸 Español

1. Clona el repositorio:
   ```bash
   git clone https://github.com/tu-usuario/Backend-Go.git
   cd backend-go
   ```

2. Crea un archivo `.env` con tus variables de entorno.

3. Instala las dependencias:
   ```bash
   go mod tidy
   ```

4. Ejecuta la aplicación:
   ```bash
   go run main.go
   ```

---

## 🧪 Tests / Pruebas

```bash
go test ./...
```

---

## 📄 Example `.env` file / Archivo `.env` de ejemplo

```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=tu_clave
DB_NAME=backend_db
JWT_SECRET=una_clave_secreta
```

> ⚠️ **Important / Importante:**  
> Don't forget to add `.env` to your `.gitignore` to avoid committing sensitive data.  
> No olvides agregar `.env` al archivo `.gitignore` para evitar subir datos sensibles.

---

## 🐳 Using Docker (optional) / Usando Docker (opcional)

```bash
docker build -t backend-go .
docker run -p 8080:8080 backend-go
```

---

## 📬 Author / Autor

**AndresGT**  
🌐 [dvandresgt.lat](https://dvandresgt.lat)  
📧 diegogiraldo506@gmail.com

---

## 📌 API Endpoints (Resumen)

### 🌐 Public

#### 🔐 Auth
- `POST /api/login` — Iniciar sesión
- `POST /api/register` — Registrar nuevo usuario

### 🔒 Private

#### 👤 Users
- `GET /api/saludo` — Obtén un mensaje con la ID del usuario actual

...