# LibraryAPI Project Plan

## Project Overview
LibraryAPI is a RESTful API written in Go that manages books and users. It uses Gin for HTTP routing, GORM for ORM, and PostgreSQL (containerized via Docker) for persistence. The program entrypoint is `main.go`, which starts the DB, runs migrations, and launches the server.

---

## Work Completed
A unified list of implemented features and configuration (no attribution split):

- Project scaffolding and Go module initialization
- Docker Compose configuration for PostgreSQL (`docker-compose.yml`)
- Database connection and pool configuration in `database/database.go`
- Migrations runner in `database/migrations/migrations.go` (auto-migrates models)

- Books
  - Model: `models/book.go` — fields: `ID`, `Name`, `Description`, `MediumPrice`, `Author`, `ImageUrl`, timestamps and soft-delete (`gorm.DeletedAt`).
  - Controller: `controllers/book_controller.go` — endpoints implemented include:
    - `GET /api/v1/books/:id` (ShowBook)
    - `GET /api/v1/books/` (ShowBooks)
    - `POST /api/v1/books/` (CreateBook)
  - Routing: registered under `/api/v1/books` in `server/routes/router.go`.
  - Migration: `database/migrations/migrations.go` runs `AutoMigrate(models.Book{})`.

- Users
  - Model: `models/user.go` — fields: `ID`, `Login`, `Password`, `CreatedAt`, `Age`, soft-delete.
  - Controller: `controllers/user_controller.go` — implemented endpoints:
    - `GET /api/v1/users/:id` (ShowUser)
    - `GET /api/v1/users/` (ShowUsers)
    - `POST /api/v1/users/` (CreateUser)
  - Routing: registered under `/api/v1/users` in `server/routes/router.go`.
  - Migration: `database/migrations/migrations.go` runs `AutoMigrate(models.User{})`.

---

## Next Steps / TODO
- Add authentication, secure password hashing (bcrypt) and login flow for `User`.
- Add update (PUT/PATCH) and delete (DELETE) endpoints for Book and User resources.
- Add input validation and better error responses (use `github.com/go-playground/validator` or Gin binding tags).
- Move database credentials and other config to environment variables (or a config file).
- Add unit and integration tests (suggest `tests/` folder); include a test Docker Compose profile if needed.
- Add API documentation (OpenAPI/Swagger or a markdown doc with examples).

---

## How to run locally (current project conventions)
1. Start PostgreSQL container

```powershell
docker-compose up db
```

2. Run the server

```powershell
go run main.go
```

Notes: DB connection string is currently hardcoded in `database/database.go` as:
`host=localhost port=25432 user=admin dbname=books sslmode=disable password=123456` — update this before deploying.

---

## References (key files)
- `main.go` — application entrypoint
- `models/book.go`, `controllers/book_controller.go` — Book resource
- `models/user.go`, `controllers/user_controller.go` — User resource
- `server/routes/router.go` — route registration
- `database/database.go`, `database/migrations/migrations.go` — DB and migrations
- `docker-compose.yml` — PostgreSQL container

---

_This document records the concrete, discoverable state of the repository. Update as features and wiring change._
