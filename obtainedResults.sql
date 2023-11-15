--  Розробити візуалізацію максимальних значень шкідливих частинок PM2.5, PM10 в розрізі областей за вказаний період часу 

select mu.title     as measured_title,
       s.city,
       max(m.value) as max_value
from measurement m
         inner join measured_unit mu on m.id_measured_unit = mu.id_measured_unit
         inner join station s on m.id_station = s.id_station
where (mu.title = 'PM2.5'
    or mu.title = 'PM10')
  and m.time between '2022-02-26' and '2023-04-05'
group by mu.title, s.city;

-- Розробити візуалізацію кількості разів, коли було зафіксовано середньодобові значення твердих частинок PM2.5, 
-- значення яких належать до шкідливого рівня на певній станції за весь час

select date_trunc('day', m.time)                              as day,
       avg(m.value)                                           as avg_value,
       count(case when ov.id_measured_unit >= '4' then 1 end) as harmful_measurements_count
from measurement m
         inner join optimal_value ov on m.id_measured_unit = ov.id_measured_unit
         inner join category c on c.id_category = ov.id_category
where ov.id_category >= '4' -- PM2.5
  and m.id_station = '0002'
group by day
order by day;

-- Розробити візуалізацію кількості вимірювань, які належать до категорій оптимальних значень для діоксиду сірки

select date_trunc('day', m.time) as day, avg(m.value) as avg_value
from measurement m
where m.id_measured_unit = '15' -- Sulfur dioxide
group by day
order by day, avg_value;

-- Розробити візуалізацію кількості вимірювань, які належать до категорій оптимальних значень для чадного газу

select date_trunc('day', m.time) as day, avg(m.value) as avg_value
from measurement m
where m.id_measured_unit = '12' -- Carbon monoxide(CO)
group by day
order by day, avg_value;