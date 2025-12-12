package models

type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"-"` // hide password in JSON
}
