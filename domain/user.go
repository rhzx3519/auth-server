package domain

type User struct {
    ID       int64  `json:"id"`
    Email    string `json:"email"`
    Password string `json:"password"`
    Nickname string `json:"nickname"`
    No       string `json:"No"`
}
