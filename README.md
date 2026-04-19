# 🇩🇪 B1 Trainer

A full-stack flashcard app for German B1 exam preparation, built with spaced repetition (SM-2 algorithm).

Built as a hobby project to demonstrate the same REST API implemented in two backend languages.

---

## Architecture

React Frontend → Spring Boot Backend (Java)
OR
React Frontend → Go Backend
↓
PostgreSQL

## Features

- 📚 Three card types — Vocabulary, Grammar, Fill-in-the-blank
- 🧠 SM-2 spaced repetition algorithm — cards scheduled by memory strength
- 📊 Dashboard — due cards, total cards
- ✏️ Add / delete cards from the UI
- 🔍 Filter by type + search
- 🔄 Two backends — swap by changing one line in the frontend

---

## Tech Stack

| Layer | Technology |
|---|---|
| Frontend | React, axios, react-router-dom |
| Backend (Java) | Spring Boot 3, JPA, Flyway, PostgreSQL |
| Backend (Go) | Go, gorilla/mux, database/sql |
| Database | PostgreSQL 16 |
| DevOps | Docker |

---

## Getting Started

### Prerequisites
- Java 21
- Node.js 20+
- Go 1.22+
- Docker Desktop

### 1. Start PostgreSQL

```bash
docker run --name b1-postgres \
  -e POSTGRES_USER=b1user \
  -e POSTGRES_PASSWORD=b1pass \
  -e POSTGRES_DB=b1trainer \
  -p 5432:5432 \
  -d postgres:16
```

### 2. Run Java Backend

```bash
cd backend
./mvnw spring-boot:run
# runs on http://localhost:8080
# Swagger UI: http://localhost:8080/swagger-ui/index.html
```

### 3. Run Go Backend

```bash
docker build -t b1trainer-go ./backend-go
docker run --name b1trainer-go \
  -p 8081:8081 \
  -e DB_HOST=host.docker.internal \
  -d b1trainer-go
# runs on http://localhost:8081
```

### 4. Run Frontend

```bash
cd frontend
npm install
npm start
# runs on http://localhost:3000
```

To switch backends change `BASE_URL` in `frontend/src/api/api.js`:

```javascript
// Java backend
const BASE_URL = 'http://localhost:8080/api';

// Go backend
const BASE_URL = 'http://localhost:8081/api';
```

---

## SM-2 Algorithm

Cards are scheduled using the SM-2 spaced repetition algorithm. After each card you rate your recall:

| Rating | Meaning | Result |
|---|---|---|
| 0 | Complete blackout | Reset to day 1 |
| 2 | Wrong but easy to recall | Reset to day 1 |
| 3 | Correct but hard | Interval grows slowly |
| 4 | Correct, good recall | Interval grows normally |
| 5 | Perfect, instant recall | Interval grows fast |

---

## Project Structure

b1-trainer/
├── backend/              # Spring Boot
│   └── src/main/java/com/b1trainer/
│       ├── card/         # Card entity, repo, controller
│       ├── progress/     # UserCardProgress + service
│       ├── session/      # Study session endpoints
│       └── algorithm/    # SM-2 implementation
├── backend-go/           # Go
│   ├── handlers/         # HTTP handlers
│   ├── models/           # Structs
│   ├── algorithm/        # SM-2 implementation
│   ├── db/               # Database connection
│   └── middleware/       # CORS
└── frontend/             # React
└── src/
├── api/          # axios calls
├── components/   # Navbar
└── pages/        # Home, Study, Cards

---

## Author

Pooran Chandrashekaraiah — [github.com/pooranc](https://github.com/pooranc)