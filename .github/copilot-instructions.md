# Copilot Instructions for LibraryAPI

## Project Overview
- **LibraryAPI** is a Go-based REST API for managing books, using Gin for HTTP routing and GORM for PostgreSQL ORM.
- The API is containerized with Docker and uses a PostgreSQL database, configured via `docker-compose.yml`.
- Main entrypoint: `main.go` initializes the database, runs migrations, and starts the server.

## Architecture & Key Components
- **Controllers** (`controllers/`): Handle HTTP requests. Example: `book_controller.go` implements CRUD for books.
- **Models** (`models/`): Define data structures. Example: `Book` struct in `book.go`.
- **Database** (`database/`): Connection logic in `database.go`, migrations in `migrations/migrations.go`.
- **Server** (`server/`): Server setup in `server.go`, routes in `routes/router.go`.
- **Routes**: All API endpoints are grouped under `/api/v1/books`.

## Developer Workflows
- **Run the API locally**:
  1. Start PostgreSQL with Docker: `docker-compose up db`
  2. Run the Go server: `go run main.go`
- **Database connection**: Hardcoded in `database.go` (host: localhost, port: 25432, user: admin, db: books, password: 123456).
- **Migrations**: Auto-run on startup via `migrations.RunMMigrations`.
- **Testing**: No test files detected; add tests in a `tests/` directory if needed.

## Patterns & Conventions
- **Gin** is used for routing; all controllers expect a `*gin.Context` parameter.
- **GORM** is used for all DB operations; models use GORM tags for schema.
- **Error handling**: Controllers return JSON error messages with HTTP 400 for invalid input or DB errors.
- **API versioning**: All endpoints are under `/api/v1`.
- **Environment variables**: DB credentials are set in `docker-compose.yml` and hardcoded in Go (consider refactoring for production).

## Integration Points
- **External dependencies**: See `go.mod` for Gin, GORM, and related libraries.
- **Docker**: PostgreSQL runs in a container; data is persisted in a named volume (`libraryapi`).

## Example: Add a New Resource
1. Create a new model in `models/`.
2. Add controller functions in `controllers/`.
3. Register routes in `server/routes/router.go`.
4. Update migrations in `database/migrations/migrations.go`.

## References
- Entrypoint: `main.go`
- Routing: `server/routes/router.go`
- DB setup: `database/database.go`, `docker-compose.yml`
- Model example: `models/book.go`
- Controller example: `controllers/book_controller.go`

---
_Review and update these instructions as the project evolves. For questions, check referenced files for implementation details._
