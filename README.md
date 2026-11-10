# 🔗 URL Shortener

A backend URL shortening service built with **Go, Gin, PostgreSQL, and Redis**. It converts long URLs into short, shareable links and redirects users to their original destinations.

The project follows a clean, layered architecture with a focus on maintainability, performance, and scalability.

## ✨ Features

- **URL Shortening** – Generate unique short links from long URLs.
- **URL Redirection** – Redirect short links to their original destinations.
- **Authentication** – Secure user authentication using JWT.
- **PostgreSQL** – Persistent storage for users and URL mappings.
- **Redis Caching** – Faster access to frequently requested URLs.
- **Clean Architecture** – Separation of handlers, services, and repositories.
- **Database Migrations** – Version-controlled database schema.
- **Docker Support** – Containerized application setup.

## 🛠️ Tech Stack

| Technology | Purpose |
|---|---|
| Go | Backend development |
| Gin | HTTP framework |
| PostgreSQL | Database |
| Redis | Caching |
| JWT | Authentication |
| pgx | PostgreSQL driver |
| Docker | Containerization |

## 🏗️ Architecture

The application follows a layered architecture:

```text
Client
  │
  ▼
Router
  │
  ▼
Handler
  │
  ▼
Service
  │
  ▼
Repository
  │
  ▼
PostgreSQL
```

Redis is used as a caching layer to reduce repeated database queries during URL redirection.

## 📁 Project Structure

```text
url-shortener/
├── cmd/
│   └── server/
│       └── main.go
├── config/
├── internal/
│   ├── handler/
│   ├── service/
│   └── repository/
├── migrations/
├── pkg/
├── .env.example
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.md
```

## 🚀 Getting Started

### Prerequisites

- Go 1.26+
- PostgreSQL
- Redis (if using caching)
- Docker (optional)

### 1. Clone the repository

```bash
git clone https://github.com/AbhayBajpai10/url-shortener.git
cd url-shortener
```

### 2. Configure environment variables

```bash
cp .env.example .env
```

Update the database credentials, JWT secret, and Redis configuration in `.env`.

### 3. Run database migrations

```bash
# Apply migrations using the project's migration setup
```

### 4. Start the server

```bash
go mod download
go run ./cmd/server
```

The server runs at:

```text
http://localhost:8080
```

## 🔄 How It Works

**1. Shorten URL**

```text
Long URL → API → Generate Short ID → PostgreSQL
                                      ↓
                                Short URL
```

**2. Redirect URL**

```text
Short URL → API → Redis Cache
                       │
                 Cache Miss?
                       │
                       ▼
                  PostgreSQL
                       │
                       ▼
                 Original URL
                       │
                       ▼
                 HTTP Redirect
```

## 🎯 Learning Outcomes

- Building REST APIs with Go and Gin.
- Implementing JWT-based authentication.
- Designing layered backend architecture.
- Working with PostgreSQL and database migrations.
- Implementing Redis caching.
- Understanding URL redirection and backend performance.

## 👨‍💻 Author

**Abhay Bajpai**

- GitHub: [@AbhayBajpai10](https://github.com/AbhayBajpai10)

---

⭐ If you find this project useful, consider giving it a star!
