# Go REST API: Users DOB and Calculated Age

A robust RESTful API built in Go using GoFiber, PostgreSQL, SQLC database wrapper, Uber Zap logging, and `go-playground/validator` input validation. 

The API calculates and returns a user's age dynamically based on their date of birth (DOB) and the server's current system date.

---

## 🚀 Features

- **Dynamic Age Calculation:** Automatically calculates users' ages based on DOB when fetching details.
- **Strict Validation:** Date formats are verified to match `YYYY-MM-DD`.
- **Database Abstraction:** DB access wrapper conforms to SQLC standard definitions.
- **Request Logging:** Middleware logs HTTP method, endpoint, response status, unique `X-Request-ID` and request duration.
- **Uber Zap Integration:** Fast, structured JSON logging.
- **Docker Support:** Ready to run inside containerized environments with `docker-compose`.

---

## 🛠️ Project Structure

```
.
├── cmd
│   └── server
│       └── main.go       # Application Entry Point
├── config
│   └── config.go        # Configuration loader (Env variables)
├── db
│   ├── migrations       # Database migrations
│   ├── queries          # SQL query definitions
│   └── sqlc             # Generated DB access code
├── internal
│   ├── handler          # HTTP request handlers (GoFiber)
│   ├── logger           # Uber Zap setup
│   ├── middleware       # Logging & Request ID Middlewares
│   ├── models           # Payloads & Validations
│   ├── repository       # DB Repository interface
│   ├── routes           # Route specifications
│   └── service          # Business logic and Age utilities
├── Dockerfile           # Multi-stage Docker deployment script
├── docker-compose.yml   # Multi-container network deployment
└── go.mod
```

---

## 🏁 How to Run

### Method 1: Docker Compose (Recommended)

Spins up a PostgreSQL instance and the Go API container automatically.

1. Ensure Docker and Docker Desktop are running on your system.
2. Run the following command from the root directory:
   ```bash
   docker-compose up --build
   ```
3. The API will listen on `http://localhost:3000`.

### Method 2: Local Setup

1. Make sure Go 1.25+ and PostgreSQL are installed on your machine.
2. Start PostgreSQL and create a database (e.g. `userdb`).
3. Set the required environment variables:
   - **Windows (PowerShell):**
     ```powershell
     $env:DATABASE_URL="postgres://<username>:<password>@localhost:5432/<database_name>?sslmode=disable"
     $env:PORT="3000"
     ```
   - **macOS/Linux:**
     ```bash
     export DATABASE_URL="postgres://<username>:<password>@localhost:5432/<database_name>?sslmode=disable"
     export PORT="3000"
     ```
4. Run the application:
   ```bash
   go run cmd/server/main.go
   ```

---

## 🧪 Running Unit Tests

Run tests targeting the service layer and age calculation logic:
```bash
go test -v ./internal/service/...
```

---

## 🔄 API Endpoints

### 1. Create User
- **Method:** `POST`
- **URL:** `/users`
- **Request Body:**
  ```json
  {
    "name": "Alice",
    "dob": "1990-05-10"
  }
  ```
- **Response (201 Created):**
  ```json
  {
    "id": 1,
    "name": "Alice",
    "dob": "1990-05-10"
  }
  ```

### 2. Get User by ID
- **Method:** `GET`
- **URL:** `/users/:id`
- **Response (200 OK):**
  ```json
  {
    "id": 1,
    "name": "Alice",
    "dob": "1990-05-10",
    "age": 36
  }
  ```

### 3. Update User
- **Method:** `PUT`
- **URL:** `/users/:id`
- **Request Body:**
  ```json
  {
    "name": "Alice Updated",
    "dob": "1991-03-15"
  }
  ```
- **Response (200 OK):**
  ```json
  {
    "id": 1,
    "name": "Alice Updated",
    "dob": "1991-03-15"
  }
  ```

### 4. Delete User
- **Method:** `DELETE`
- **URL:** `/users/:id`
- **Response:** `204 No Content`

### 5. List All Users (with Pagination)
- **Method:** `GET`
- **URL:** `/users?page=1&limit=10`
- **Response (200 OK):**
  ```json
  [
    {
      "id": 1,
      "name": "Alice Updated",
      "dob": "1991-03-15",
      "age": 35
    }
  ]
  ```
