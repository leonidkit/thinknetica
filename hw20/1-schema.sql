DROP TABLE IF EXISTS film_actor;
DROP TABLE IF EXISTS film_producer;
DROP TABLE IF EXISTS film;

DROP TABLE IF EXISTS studio;

DROP TABLE IF EXISTS actor;
DROP TABLE IF EXISTS producer;


-- Таблица студия
-- Столбцы: название
-- (у одной студии может быть много фильмов)
CREATE TABLE studio (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL UNIQUE
);

-- Таблица актер
-- Столбцы: имя, дата рождения
CREATE TABLE actor(
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    born DATE NOT NULL
    -- born BIGINT NOT NULL
);

-- Таблица режиссер
-- Столбцы: имя, дата рождения
CREATE TABLE producer(
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    born DATE NOT NULL
    -- born BIGINT NOT NULL
);

-- Определение типа для возрастного рейтинга
DROP TYPE IF EXISTS PG;
CREATE TYPE PG AS ENUM('PG-10', 'PG-13', 'PG-18');

-- Таблица фильм
-- Отношение к студии: один к одному
-- Отношение к актеру: один ко многим
-- Отношение к режиссеру: один ко многим
-- Столбцы: название (уникально в разрезе одного года), год выхода (>= 1800), актёры, режиссёры, сборы, рейтинг (PG-10, PG-13, PG-18)
CREATE TABLE film (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    release_date DATE NOT NULL CHECK (release_date >= '1800-01-01'::DATE),
    -- release_date BIGINT NOT NULL CHECK (to_timestamp(release_date)::DATE <= '1800-01-01'::DATE),
    box_office MONEY NOT NULL DEFAULT 0,
    rating PG NOT NULL,
    studio_id BIGINT NOT NULL REFERENCES studio(id) ON DELETE CASCADE ON UPDATE CASCADE DEFAULT 0
);

-- Функция для проверки полей фильма
DROP FUNCTION IF EXISTS check_film;
CREATE OR REPLACE FUNCTION check_film()
    RETURNS trigger AS $$
DECLARE
    year int;
    count int;
BEGIN
    SELECT EXTRACT(YEAR FROM NEW.release_date) INTO year;
    SELECT count(id) INTO count FROM film WHERE title = NEW.title AND EXTRACT(YEAR FROM release_date) = year;
    
    IF count != 0 
        THEN RAISE EXCEPTION 'Invalid film title. Already exists in the % year', year;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Создание тригера с функцией для проверки полей фильма
DROP TRIGGER IF EXISTS check_film ON film;
CREATE TRIGGER check_film BEFORE INSERT OR UPDATE ON film
    FOR EACH ROW EXECUTE PROCEDURE check_film();

-- Связь между фильмами и актерами
-- (у одного фильма может быть много актеров)
CREATE TABLE film_actor (
    id BIGSERIAL PRIMARY KEY,
    film_id BIGINT NOT NULL REFERENCES film,
    actor_id BIGINT NOT NULL REFERENCES actor,
    UNIQUE (film_id, actor_id) 
);

-- Связь между фильмами и режиссерами
-- (у одного фильма может быть много режиссеров)
CREATE TABLE film_producer (
    id BIGSERIAL PRIMARY KEY,
    film_id BIGINT NOT NULL REFERENCES film,
    producer_id BIGINT NOT NULL REFERENCES producer,
    UNIQUE (film_id, producer_id) 
);