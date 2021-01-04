DROP TABLE IF EXISTS FILM_ACTOR;
DROP TABLE IF EXISTS FILM_DIRECTOR;
DROP TABLE IF EXISTS FILM;

DROP TABLE IF EXISTS PRODUCTION_COMPANY;

DROP TABLE IF EXISTS ACTOR;
DROP TABLE IF EXISTS DIRECTOR;


-- Таблица студия
-- Столбцы: название
-- (у одной студии может быть много фильмов)
CREATE TABLE PRODUCTION_COMPANY (
    ID BIGSERIAL PRIMARY KEY,
    TITLE TEXT NOT NULL UNIQUE
);

-- Таблица актер
-- Столбцы: имя, дата рождения
CREATE TABLE ACTOR(
    ID BIGSERIAL PRIMARY KEY,
    NAME TEXT NOT NULL,
    BORN DATE NOT NULL
);

-- Таблица режиссер
-- Столбцы: имя, дата рождения
CREATE TABLE DIRECTOR(
    ID BIGSERIAL PRIMARY KEY,
    NAME TEXT NOT NULL,
    BORN DATE NOT NULL
);

-- Определение типа для возрастного рейтинга
DROP TYPE IF EXISTS PG;
CREATE TYPE PG AS ENUM('PG-10', 'PG-13', 'PG-18');

-- Таблица фильм
-- Отношение к студии: один к одному
-- Отношение к актеру: один ко многим
-- Отношение к режиссеру: один ко многим
-- Столбцы: название (уникально в разрезе одного года), год выхода (>= 1800), актёры, режиссёры, сборы, рейтинг (PG-10, PG-13, PG-18)
CREATE TABLE FILM (
    ID BIGSERIAL PRIMARY KEY,
    TITLE TEXT NOT NULL,
    RELEASE_DATE DATE NOT NULL,
    BOX_OFFICE NUMERIC NOT NULL DEFAULT 0 CHECK (BOX_OFFICE >= 0),
    RATING PG NOT NULL,
    PRODUCTION_COMPANY_ID BIGINT NOT NULL REFERENCES PRODUCTION_COMPANY(ID) ON DELETE CASCADE ON UPDATE CASCADE DEFAULT 0
);

-- Функция для проверки полей фильма
DROP FUNCTION IF EXISTS check_film;
CREATE OR REPLACE FUNCTION check_film()
    RETURNS trigger AS $$
DECLARE
    year int;
    count int;
BEGIN
    IF NEW.RELEASE_DATE <= '1800-01-01'::DATE 
        THEN RAISE EXCEPTION 'Invalid release date';
    END IF;


    SELECT EXTRACT(YEAR FROM NEW.RELEASE_DATE) INTO year;
    SELECT count(id) INTO count FROM FILM WHERE TITLE = NEW.TITLE AND EXTRACT(YEAR FROM RELEASE_DATE) = year;
    
    IF count != 0 
        THEN RAISE EXCEPTION 'Invalid film title. Already exists in the % year', year;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Создание тригера с функцией для проверки полей фильма
DROP TRIGGER IF EXISTS check_film ON FILM;
CREATE TRIGGER check_film BEFORE INSERT OR UPDATE ON FILM
    FOR EACH ROW EXECUTE PROCEDURE check_film();

-- Связь между фильмами и актерами
-- (у одного фильма может быть много актеров)
CREATE TABLE FILM_ACTOR (
    ID BIGSERIAL PRIMARY KEY,
    FILM_ID BIGINT NOT NULL REFERENCES FILM,
    ACTOR_ID BIGINT NOT NULL REFERENCES ACTOR,
    UNIQUE (FILM_ID, ACTOR_ID) 
);

-- Связь между фильмами и режиссерами
-- (у одного фильма может быть много режиссеров)
CREATE TABLE FILM_DIRECTOR (
    ID BIGSERIAL PRIMARY KEY,
    FILM_ID BIGINT NOT NULL REFERENCES FILM,
    DIRECTOR_ID BIGINT NOT NULL REFERENCES DIRECTOR,
    UNIQUE (FILM_ID, DIRECTOR_ID) 
);