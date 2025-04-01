
# Simple REST API in Go using net/http

This example demonstrates how to implement a basic REST API in Go using the standard `net/http` package.

---

## 📦 Features

- POST /user — Add a new user
- GET /user — Retrieve all users
- In-memory data storage
- Thread-safe with mutex

---

## 📝 Code Example

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "sync"
)

// User represents a simple user object
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var (
    users  []User
    mu     sync.Mutex
    nextID = 1
)

func main() {
    http.HandleFunc("/user", userHandler)

    fmt.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// userHandler handles GET and POST /user
func userHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    switch r.Method {
    case http.MethodGet:
        mu.Lock()
        defer mu.Unlock()
        json.NewEncoder(w).Encode(users)

    case http.MethodPost:
        var user User
        err := json.NewDecoder(r.Body).Decode(&user)
        if err != nil {
            http.Error(w, "Invalid input", http.StatusBadRequest)
            return
        }

        mu.Lock()
        user.ID = nextID
        nextID++
        users = append(users, user)
        mu.Unlock()

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(user)

    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
```

---

## 🧪 Example Requests

### POST /user

```bash
curl -X POST http://localhost:8080/user \
  -H "Content-Type: application/json" \
  -d '{"name": "Alice"}'
```

### GET /user

```bash
curl http://localhost:8080/user
```

---

## ✅ Notes

- Run with:
  ```bash
  go run main.go
  ```

- Access the API on: `http://localhost:8080/user`
