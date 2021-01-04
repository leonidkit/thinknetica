-- выборка фильмов с названием студии
SELECT 
    f.TITLE,
    pc.TITLE  
FROM 
    FILM f 
    JOIN 
        PRODUCTION_COMPANY pc 
        ON f.PRODUCTION_COMPANY_ID = pc.ID;

-- выборка фильмов для некоторого актёра
SELECT
   f.TITLE,
   a.NAME 
FROM
   FILM_ACTOR fa 
   JOIN
      FILM f 
      ON f.ID = fa.FILM_ID 
   JOIN
      ACTOR a 
      ON a.ID = fa.ACTOR_ID 
WHERE
   a.NAME = 'Mills, Adria J.';

-- подсчёт фильмов для некоторого режиссёра
SELECT
   count(1)
FROM
   FILM_DIRECTOR fd
   JOIN
      DIRECTOR d
      ON d.ID = fd.DIRECTOR_ID 
WHERE
   d.NAME = 'Warner, Sasha R.';

-- выборка фильмов для нескольких режиссёров из списка (подзапрос)
-- непонятна формулировка
SELECT
   f.TITLE,
   a.NAME 
FROM
   FILM_ACTOR fa 
   JOIN
      FILM f 
      ON f.ID = fa.FILM_ID 
   JOIN
      ACTOR a 
      ON a.ID = fa.ACTOR_ID 
WHERE
   a.NAME IN ('Warner, Sasha R.', 'Mills, Adria J.');

-- подсчёт количества фильмов для актёра
SELECT
   COUNT(1)
FROM
   FILM_ACTOR fa 
   JOIN
      ACTOR a 
      ON a.ID = fa.ACTOR_ID 
WHERE
   a.NAME = 'Mills, Adria J.';
   
-- выборка актёров и режиссёров, участвовавших более чем в 2 фильмах
SELECT
   a.NAME,
   count(1) 
FROM
   FILM_ACTOR fa 
   JOIN
      ACTOR a 
      ON fa.ACTOR_ID = a.ID 
GROUP BY
(a.NAME) 
HAVING
   count(1) > 2 
UNION ALL
SELECT
   d.NAME,
   count(1) 
FROM
   FILM_DIRECTOR fd 
   JOIN
      DIRECTOR d 
      ON fd.DIRECTOR_ID = d.ID 
GROUP BY
(d.NAME) 
HAVING
   count(1) > 2

-- подсчёт количества фильмов со сборами больше 1000
SELECT count(1) FROM FILM WHERE BOX_OFFICE > 1000

-- подсчитать количество режиссёров, фильмы которых собрали больше 1000
SELECT
   count(1) 
FROM
   FILM f 
   JOIN
      FILM_DIRECTOR fd 
      ON fd.FILM_ID = f.ID 
WHERE
   BOX_OFFICE > 1000

-- выборка различных фамилий актёров
SELECT DISTINCT a.NAME FROM ACTOR a

-- подсчёт количества фильмов, имеющих дубли по названию
SELECT
   COUNT(1) 
FROM
   (
      SELECT
         TITLE
      FROM
         FILM 
      GROUP BY
         TITLE 
      HAVING
         COUNT(1) > 1
   )
   as a