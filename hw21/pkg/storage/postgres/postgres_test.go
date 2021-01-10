package postgres

import (
	"database/sql"
	"hw21/pkg/storage"
	"log"
	"os"
	"testing"
)

var (
	strg  = &DB{}
	films = []storage.Film{
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

func TestMain(m *testing.M) {
	var err error
	strg, err = New("127.0.0.1", "5432", "admin", "admin", "store", false)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	defer strg.Close()

	_, err = strg.db.Exec("DELETE FROM film")
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	_, err = strg.db.Exec("ALTER SEQUENCE film_id_seq RESTART WITH 1")
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestDB_AddFilms(t *testing.T) {
	err := strg.AddFilms(films)
	if err != nil {
		log.Fatalf("обнаружена ошибка: %+v\n", err)
	}

	// проверяем, что все необходимые строки вставлены
	rows, err := strg.db.Query("SELECT * FROM film")
	for rows.Next() {
		var film storage.Film

		rows.Scan(
			&film.ID,
			&film.Title,
			&film.ReleaseDate,
			&film.BoxOffice,
			&film.Rating,
			&film.StudioID,
		)

		found := false
		for _, f := range films {
			if f.ID == film.ID {
				found = true
			}
		}
		if !found {
			log.Fatalf("ошибка вставки записей в бд: среди добавленных записей не было записи с ID == %d", film.ID)
		}
	}

	// попробовать добавить пустой список
	err = strg.AddFilms([]storage.Film{})
	if err == nil {
		log.Fatal("ожидалась ошибка, но не была получена")
	}
}

func TestDB_DeleteFilm(t *testing.T) {
	err := strg.DeleteFilm(films[0])
	if err != nil {
		log.Fatalf("обнаружена ошибка: %+v\n", err)
	}

	row := strg.db.QueryRow("SELECT * FROM film WHERE id = $1", films[0].ID)
	err = row.Scan()
	if err != sql.ErrNoRows {
		log.Fatalf("ожидалось, что запись с ID = %d будет удалена", films[0].ID)
	}

	err = strg.DeleteFilm(storage.Film{ID: 0})
	if err == nil {
		log.Fatal("ожидалась ошибка, но не была получена")
	}
}

func TestDB_UpdateFilm(t *testing.T) {
	filmNew := films[1]
	filmNew.Title = "Большой куш 2"

	err := strg.UpdateFilm(filmNew)
	if err != nil {
		log.Fatalf("обнаружена ошибка: %+v\n", err)
	}

	var f storage.Film
	row := strg.db.QueryRow("SELECT * FROM film WHERE id = $1", films[1].ID)
	err = row.Scan(
		&f.ID,
		&f.Title,
		&f.ReleaseDate,
		&f.BoxOffice,
		&f.Rating,
		&f.StudioID,
	)
	if err != nil {
		log.Fatalf("обнаружена ошибка: %+v\n", err)
	}

	if f.Title != filmNew.Title {
		log.Fatalf("ожидалась запись с Title = %s, а получена Title = %s", filmNew.Title, films[1].Title)
	}

	err = strg.UpdateFilm(storage.Film{ID: 0})
	if err == nil {
		log.Fatal("ожидалась ошибка, но не была получена")
	}
}

func TestDB_Films(t *testing.T) {
	res, err := strg.Films(0)
	if err != nil {
		log.Fatalf("обнаружена ошибка: %+v\n", err)
	}

	found := 0
	for _, f1 := range res {
		for _, f2 := range films[1:] {
			if f1.ID == f2.ID {
				found++
			}
		}
	}

	if found != len(films[1:]) {
		log.Fatalf("ожидалось %d записей, но найдено %d", len(films[1:]), found)
	}
}
