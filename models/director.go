package models

type Celebrity struct {
	ID          uint     `gorm:"primaryKey"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	Is_actor    bool     `json:"is_actor"`
	Is_producer bool     `json:"is_producer"`
	TotalMovies int      `json:"total_movies"`
	Movie       []*Movie `gorm:"many2many:movie_celebrities;"`
}
