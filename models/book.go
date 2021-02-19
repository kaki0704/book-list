package models

// Book is ...
type Book struct {
	ID     int    `gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}
