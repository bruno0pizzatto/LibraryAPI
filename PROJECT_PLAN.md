# LibraryAPI Project Plan

## Project Overview
LibraryAPI is a RESTful API built with Go, Gin, and GORM, designed to manage books and users. The backend uses PostgreSQL, containerized via Docker.

---

## Initial Setup (Solo)
- Initialized Go module and project structure
- Created Docker Compose for PostgreSQL
- Implemented Book model, controller, and endpoints
- Set up database connection and migration logic
- Configured Gin server and routing

---

## AI-Assisted Additions
- Generated and documented `.github/copilot-instructions.md` for AI agent productivity
- Added User model (`models/user.go`) with fields: login, password, creation date, age
- Created User controller (`controllers/user_controller.go`) with endpoints:
  - `GET /api/v1/users/:id` (ShowUser)
  - `GET /api/v1/users/` (ShowUsers)
  - `POST /api/v1/users/` (CreateUser)
- Updated migration logic to include User table
- Registered User endpoints in router (`server/routes/router.go`)

---

## Next Steps / TODO
- Add authentication and password hashing for User
- Implement update and delete endpoints for Book and User
- Add automated tests (suggested: create `tests/` directory)
- Refactor DB credentials to use environment variables
- Improve error handling and validation
- Document API endpoints and usage examples

---

## References
- Entrypoint: `main.go`
- Book resource: `models/book.go`, `controllers/book_controller.go`
- User resource: `models/user.go`, `controllers/user_controller.go`
- Routing: `server/routes/router.go`
- Migrations: `database/migrations/migrations.go`
- DB setup: `database/database.go`, `docker-compose.yml`

---

_This plan documents all major steps taken solo and with GitHub Copilot. Update as the project evolves._
