package handlers

import (
    "encoding/json"
    "net/http"

    "golang.org/x/crypto/bcrypt"
    "myapp/database"
    "myapp/models"
)

type RegisterRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
        return
    }

    var req RegisterRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }

    user := models.User{
        Username: req.Username,
        Password: string(hashedPassword),
    }

    stmt, err := database.DB.Prepare("INSERT INTO users (username, password) VALUES (?, ?)")
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }
    defer stmt.Close()

    _, err = stmt.Exec(user.Username, user.Password)
    if err != nil {
        http.Error(w, "Username already exists", http.StatusBadRequest)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("User registered successfully"))
}
