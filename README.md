# Digital CV

A full-stack digital CV web app with a Go backend (PostgreSQL) and a Next.js frontend.

## Project Structure

- `backend/` - Go API server
  - `cmd/api/main.go` - server entrypoint
  - `internal/db` - postgres connection + CRUD repository
  - `internal/handlers` - HTTP handlers
  - `internal/services` - business logic
  - `internal/models` - data models
- `db/` - database initialization schema (`init.sql`)
- `frontend/` - Next.js app (app router)
  - `app/page.tsx` - homepage
  - `public/` - static files

## Requirements

- Go 1.22+ (or latest stable)
- Docker + docker-compose (optional)
- PostgreSQL (via docker compose or installed locally)
- Node.js 18+ and npm/yarn/pnpm

## Development

### 1) Start database

```bash
cd /path/to/digital-cv
# using docker compose
docker-compose up -d
```

Database is configured by `db/init.sql`.

### 2) Start backend

```bash
cd backend
go run ./cmd/api
```

API listens on `http://localhost:8080` (default in code)

### 3) Start frontend

```bash
cd frontend
npm install
npm run dev
```

Open `http://localhost:3000`.

## API

Expected endpoints (verify in backend code):

- `GET /cv` - read CV data
- `POST /cv` - create CV
- `PUT /cv/:id` - update CV
- `DELETE /cv/:id` - delete CV

## Tests

Backend:

```bash
cd backend
go test ./...
```

Frontend:

```bash
cd frontend
npm test
```

## Environment

- `BACKEND_DB_HOST` / `BACKEND_DB_USER` / `BACKEND_DB_PASSWORD` / `BACKEND_DB_NAME` (from `docker-compose.yml`)
- `NEXT_PUBLIC_API_URL` (if needed) default: `http://localhost:8080`

## Deploy

- Backend: build Go binary and deploy in container or host.
- Frontend: deploy with Vercel or any static host supporting Next.js.

## Notes

This project is a personal digital curriculum vitae app in progress. Update `frontend/app/page.tsx` and backend packages to add content and features.