package postgres

import (
	"context"
	"fmt"
	"log"

	"hw21/pkg/storage"

	pgx "github.com/jackc/pgx/v4"
)

type DB struct {
	conn *pgx.Conn
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

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	pgDB := &DB{
		conn: conn,
	}

	return pgDB, nil
}

func (p *DB) AddFilms(films []storage.Film) error {
	ctx := context.Background()

	if len(films) == 0 {
		return fmt.Errorf("пустой список фильмов")
	}

	tx, err := p.conn.Begin(ctx)
	if err != nil {
		log.Printf("%+v\n", err)
		return err
	}
	defer tx.Rollback(ctx)

	batch := new(pgx.Batch)

	for _, film := range films {
		batch.Queue("INSERT INTO film(title,release_date,box_office,rating,studio_id) VALUES ($1, $2, $3, $4, $5)",
			film.Title,
			film.ReleaseDate,
			film.BoxOffice,
			film.Rating,
			film.StudioID,
		)
	}

	res := tx.SendBatch(ctx, batch)
	if err = res.Close(); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (p *DB) DeleteFilm(f storage.Film) error {
	_, err := p.conn.Exec(context.Background(), "DELETE FROM film WHERE id = $1", f.ID)
	return err
}

func (p *DB) UpdateFilm(f storage.Film) error {
	_, err := p.conn.Exec(context.Background(), `UPDATE film SET title=$2,release_date=$3,box_office=$4,rating=$5,studio_id=$6 WHERE id = $1`,
		f.ID,
		f.Title,
		f.ReleaseDate,
		f.BoxOffice,
		f.Rating,
		f.StudioID,
	)
	return err
}

func (p *DB) Films(id int64) ([]storage.Film, error) {
	var films []storage.Film
	var rows pgx.Rows
	var err error
	var ctx context.Context = context.Background()

	if id != 0 {
		sqlStatement := "SELECT * FROM film WHERE id = $1"
		rows, err = p.conn.Query(ctx, sqlStatement, id)
		if err != nil && err != pgx.ErrNoRows {
			return films, err
		}
	} else {
		sqlStatement := "SELECT * FROM film"
		rows, err = p.conn.Query(ctx, sqlStatement)
		if err != nil && err != pgx.ErrNoRows {
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

	if err = rows.Err(); err != nil {
		return films, err
	}

	return films, nil
}

func (p *DB) Close() {
	p.conn.Close(context.Background())
}
