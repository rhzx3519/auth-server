package domain

import "time"

type User struct {
    ID        int64     `json:"id"`
    Email     string    `json:"email"`
    Password  string    `json:"password"`
    Nickname  string    `json:"nickname"`
    No        string    `json:"no"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
