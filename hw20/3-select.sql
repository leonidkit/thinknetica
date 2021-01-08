-- выборка фильмов с названием студии
SELECT 
    f.title,
    s.title  
FROM 
    film f 
    JOIN 
        studio s 
        ON f.studio_id = s.id;

-- выборка фильмов для некоторого актёра
SELECT
   f.title,
   a.name 
FROM
   film_actor fa 
   JOIN
      film f 
      ON f.id = fa.fim_id 
   JOIN
      actor a 
      ON a.id = fa.actor_id 
WHERE
   a.name = 'Mills, Adria J.';

-- подсчёт фильмов для некоторого режиссёра
SELECT
   count(1)
FROM
   film_producer fp
   JOIN
      producer d
      ON d.id = fp.producer_id 
WHERE
   d.name = 'Warner, Sasha R.';

-- выборка фильмов для нескольких режиссёров из списка (подзапрос)
-- непонятна формулировка
SELECT
   f.title,
   p.name 
FROM
   film_producer fp 
   JOIN
      film f 
      ON f.id = fp.film_id 
   JOIN
      producer p
      ON p.id = fp.actor_id
WHERE
   a.name IN ('Warner, Sasha R.', 'Mills, Adria J.');



-- подсчёт количества фильмов для актёра
SELECT
   COUNT(1)
FROM
   film_actor fa 
   JOIN
      actor a 
      ON a.id = fa.actor_id 
WHERE
   a.name = 'Mills, Adria J.';
   
-- выборка актёров и режиссёров, участвовавших более чем в 2 фильмах
SELECT
   a.name,
   count(1) 
FROM
   film_actor fa 
   JOIN
      actor a 
      ON fa.actor_id = a.id 
GROUP BY
(a.name) 
HAVING
   count(1) > 2 
UNION ALL
SELECT
   d.name,
   count(1) 
FROM
   film_producer fp 
   JOIN
      producer d 
      ON fp.producer_id = d.id 
GROUP BY
(d.name) 
HAVING
   count(1) > 2

-- подсчёт количества фильмов со сборами больше 1000
SELECT count(1) FROM film WHERE box_office > 1000

-- подсчитать количество режиссёров, фильмы которых собрали больше 1000
SELECT
   count(1) 
FROM
   film f 
   JOIN
      film_producer fp 
      ON fp.fim_id = f.id 
WHERE
   box_office > 1000

-- выборка различных фамилий актёров
SELECT DISTINCT a.name FROM actor a

-- подсчёт количества фильмов, имеющих дубли по названию
SELECT
   COUNT(1) 
FROM
   (
      SELECT
         title
      FROM
         film 
      GROUP BY
         title 
      HAVING
         COUNT(1) > 1
   )
   as a