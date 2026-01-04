# Go Fiber Minimal - Clean Architecture

[![Go Version](https://img.shields.io/badge/Go-1.25.3-00ADD8?style=flat-square&logo=go)](https://golang.org/)
[![Framework](https://img.shields.io/badge/Framework-Fiber_v2-ef333f?style=flat-square)](https://gofiber.io/)
[![Architecture](https://img.shields.io/badge/Architecture-Clean_Architecture-blue?style=flat-square)](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

A minimal backend boilerplate built with **Go** and **Fiber v2** framework, following _Clean Architecture_ principles. Engineered for scalability, testability, and high performance.

## 🚀 Key Features

- **Authentication System**: Registration, Login, and Logout using JWT.
- **Security**: Robust password hashing with Bcrypt and built-in CORS protection.
- **Authorization**: Role-Based Access Control (RBAC) and granular permission management.
- **REST Client Ready**: Includes `.http` files for direct API testing within VS Code.
- **Multi-language (i18n)**: Native support for Indonesian (ID) and English (EN) translations.
- **Premium Logging**: Integrated application and database (GORM) logging with clean, readable formatting.
- **Hot Reload**: Fast development workflow supported by `air`.
- **Validation**: Centralized request validation integrated with the translation system.

## 🛠️ Tech Stack

- **Language**: [Go (Golang)](https://golang.org/)
- **Web Framework**: [Fiber v2](https://gofiber.io/)
- **ORM**: [GORM](https://gorm.io/)
- **Database**: PostgreSQL
- **Auth**: JWT (v5) & Bcrypt
- **Hot Reload**: [Air](https://github.com/cosmtrek/air)

## 📁 Project Structure

```text
.
├── app
│   ├── http
│   │   ├── controller # API endpoint logic
│   │   ├── request    # Input validation schemas
│   │   └── resource   # API response formatting (Transformers)
│   └── middleware     # Route protection (Auth, Cors, etc.)
├── bootstrap          # Application initialization
├── config             # Environment and Database configuration
├── database
│   └── entity         # GORM Models / Database Tables
├── lang               # Translation files (ID/EN) and Translator core
├── lib                # Third-party wrappers (DB, JWT, Logger, Validator)
├── logs               # Daily log files (Application & SQL)
├── rest_client        # API testing via REST Client extension (.http)
├── route              # Endpoint definitions (API & Web)
├── service            # Business logic layer
└── util               # Helper functions (Hashing, Converters, etc.)
```

## ⚙️ Getting Started

1. **Clone the Repository**

   ```bash
   git clone <repository-url>
   cd go-fiber-minimal
   ```

2. **Setup Environment**
   Create a `.env` file in the root directory and configure your variables:

   ```env
   APP_PORT=4000
   APP_ENV=local
   APP_LOCALE=en
   APP_SECRET=your_jwt_secret_key
   APP_URL=localhost:4000

   APP_DB_HOST=localhost
   APP_DB_PORT=5432
   APP_DB_NAME=go_fiber_db
   APP_DB_USER=postgres
   APP_DB_PASS=password
   APP_DB_SCHEMA=public
   APP_DB_SSLMODE=disable
   APP_DB_TimeZone=Asia/Jakarta
   ```

3. **Install Dependencies**
   ```bash
   go mod tidy
   ```

## 🏃 Running the Application

### Development Mode (Hot Reload)

Ensure you have [Air](https://github.com/cosmtrek/air) installed, then run:

```bash
air
```

### Standard Mode

```bash
go run main.go
```

## 📡 API Testing

You can use the **REST Client** extension in VS Code to test endpoints. Open the following file:

- `rest_client/auth.rest.http`

Make sure to adjust the `@HOST` variable inside the file to match your server address.

---

Made with ❤️ by Fys
