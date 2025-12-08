package models

type Milktea struct {
    ID        int     `json:"id"`
    Name      string  `json:"name"`
    Price     float64 `json:"price"`
    CreatedAt string  `json:"created_at"`
}