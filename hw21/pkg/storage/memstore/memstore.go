package memstore

import (
	"hw21/pkg/storage"
	"sync"
)

var (
	data = []storage.Film{
		{
			ID:          1,
			Title:       "Унесенные ветром",
			ReleaseDate: 1610143502,
			BoxOffice:   12312312.123,
			Rating:      "PG-13",
			StudioID:    1,
		},
		{
			ID:          2,
			Title:       "Большой куш",
			ReleaseDate: 1610143502,
			BoxOffice:   12312312.123,
			Rating:      "PG-18",
			StudioID:    2,
		},
		{
			ID:          3,
			Title:       "Старикам здесь не место",
			ReleaseDate: 1610143502,
			BoxOffice:   12312312.123,
			Rating:      "PG-18",
			StudioID:    3,
		},
		{
			ID:          4,
			Title:       "Семейка Крудс",
			ReleaseDate: 1610143502,
			BoxOffice:   12312312.123,
			Rating:      "PG-13",
			StudioID:    1,
		},
	}
)

type DB struct {
	mux   *sync.Mutex
	films []storage.Film
}

func New() *DB {
	return &DB{
		films: data,
		mux:   new(sync.Mutex),
	}
}

func (d *DB) AddFilms(f []storage.Film) error {
	d.mux.Lock()
	d.films = append(d.films, f...)
	d.mux.Unlock()
	return nil
}

func (d *DB) DeleteFilm(f storage.Film) error {
	d.mux.Lock()
	defer d.mux.Unlock()

	for idx, film := range d.films {
		if f.ID == film.ID {
			d.films[idx] = d.films[len(d.films)-1]
			d.films[len(d.films)-1] = storage.Film{}
			d.films = d.films[:len(d.films)-1]
		}
	}
	return nil
}

func (d *DB) UpdateFilm(f storage.Film) error {
	d.mux.Lock()
	defer d.mux.Unlock()

	for idx, f := range d.films {
		if f.ID == f.ID {
			d.films[idx] = f
		}
	}
	return nil
}

func (d *DB) Films(id int64) ([]storage.Film, error) {
	d.mux.Lock()
	defer d.mux.Unlock()

	if id != 0 {
		for _, f := range d.films {
			if f.ID == id {
				return []storage.Film{f}, nil
			}
		}
	}
	return d.films, nil
}
