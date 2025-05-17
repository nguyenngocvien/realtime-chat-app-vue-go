
# Realtime Chat App

## Technology used

- Golang (mux, GORM, websocket)
- Vue.js (Vite)
- PostgreSQL
- Docker & Docker Compose

---

## Project structure

```
/
├── backend/           # source code backend (Go)
├── frontend/          # source frontend (Vue)
├── docker-compose.yml # Docker Compose configure
```
---

## Require

- Docker & Docker Compose are installed on your machine.
- Ports 3000, 8080 and 5432 are not occupied

---

## Guide

1. Clone repository:

```bash
git clone https://github.com/yourusername/realtime-chat-app.git
cd realtime-chat-app
```

2. Run the entire service using Docker Compose:

```bash
docker-compose up --build
```

3. Access:

Open browser access http://localhost:3000

4. Backend API run at:

http://localhost:8080

---

## Services in Docker Compose

- `db`: PostgreSQL database, port 5432, persistent data storage by volume
- `backend`: Golang API server, using mux, GORM, WebSocket, port 8080
- `frontend`: Vue.js dev server, port 3000

---

## How to work with WebSocket

Backend uses WebSocket to send and receive realtime messages between clients.

Frontend connects WebSocket to backend and processes displaying messages as soon as new events occur.

---

## GORM & Database Migration

Backend uses GORM to map Golang struct to PostgreSQL table and automatically migrates on startup.

---

## Contact

If you have any questions or comments, please contact:

- Github: https://nguyenngocvien.github.io

---
