package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"hw21/pkg/storage"

	_ "github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

func New(host, port, user, password, dbname string, sslmode bool) (*DB, error) {
	sslmodeStr := "disable"
	if sslmode {
		sslmodeStr = "enable"
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host,
		port,
		user,
		password,
		dbname,
		sslmodeStr,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	pgDB := &DB{
		db: db,
	}

	return pgDB, nil
}

func (p *DB) AddFilms(films []storage.Film) error {
	if len(films) == 0 {
		return fmt.Errorf("пустой список фильмов")
	}

	stmt, err := p.db.Prepare("INSERT INTO film(title,release_date,box_office,rating,studio_id) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		log.Printf("%+v\n", err)
		return err
	}
	defer stmt.Close()

	for _, film := range films {
		_, err := stmt.Exec(film.Title, film.ReleaseDate, film.BoxOffice, film.Rating, film.StudioID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *DB) DeleteFilm(f storage.Film) error {
	if f.ID == 0 {
		return fmt.Errorf("отсутствует ID записи для удаления")
	}

	_, err := p.db.Exec("DELETE FROM film WHERE id = $1", f.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *DB) UpdateFilm(f storage.Film) error {
	if f.ID == 0 {
		return fmt.Errorf("отсутствует ID записи для обновления")
	}

	_, err := p.db.Exec(`UPDATE film SET title=$2,release_date=$3,box_office=$4,rating=$5,studio_id=$6 WHERE id = $1`,
		f.ID,
		f.Title,
		f.ReleaseDate,
		f.BoxOffice,
		f.Rating,
		f.StudioID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (p *DB) Films(id int64) ([]storage.Film, error) {
	var films []storage.Film
	var rows *sql.Rows
	var err error

	if id != 0 {
		sqlStatement := "SELECT * FROM film WHERE id = $1"
		rows, err = p.db.Query(sqlStatement, id)
		if err != nil {
			return films, err
		}
	} else {
		sqlStatement := "SELECT * FROM film"
		rows, err = p.db.Query(sqlStatement)
		if err != nil {
			return films, err
		}
	}

	for rows.Next() {
		film := storage.Film{}
		err = rows.Scan(
			&film.ID,
			&film.Title,
			&film.ReleaseDate,
			&film.BoxOffice,
			&film.Rating,
			&film.StudioID,
		)
		if err != nil {
			return films, err
		}

		films = append(films, film)
	}

	return films, nil
}

func (p *DB) Close() {
	p.db.Close()
}
