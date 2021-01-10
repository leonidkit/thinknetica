package storage

type Film struct {
	ID          int64   `json:"id,omitempty"`
	Title       string  `json:"title,omitempty"`
	ReleaseDate int64   `json:"release_date,omitempty"`
	BoxOffice   float64 `json:"box_office,omitempty"`
	Rating      string  `json:"rating,omitempty"`
	StudioID    int64   `json:"studio_id,omitempty"`
}

type Interface interface {
	AddFilms([]Film) error
	DeleteFilm(Film) error
	UpdateFilm(Film) error
	Films(int64) ([]Film, error) // если id == 0 то вернутся все фильмы из хранилища
}
