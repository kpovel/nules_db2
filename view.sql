create index idx_measurement_station_time on measurement (id_station, time);

create materialized view mv_first_measurements as
select distinct on (id_station, m.id_measured_unit) id_station,
                                                    mu.id_measured_unit,
                                                    first_value(time) over (w_asc) as first_measured
from measurement m
         left join measured_unit mu on mu.id_measured_unit = m.id_measured_unit
window w_asc as (partition by id_station order by time)
order by id_station;

create materialized view mv_last_measurements as
with last_measurements as (select distinct on (id_measured_unit, id_station) id_station, id_measured_unit, value
                           from measurement
                           order by id_measured_unit, id_station, time)

select lm.id_station,
       m.id_measured_unit,
       m.title,
       m.unit,
       lm.value
from measured_unit m
         inner join last_measurements lm on m.id_measured_unit = lm.id_measured_unit
order by id_station;

create view connected_stations as
select distinct on (station.id_station, mv_last_measurements.id_measured_unit) station.id_station,
                                                                               station.city,
                                                                               station.name,
                                                                               mv_first_measurements.first_measured enbled_from,
                                                                               mv_last_measurements.title           measured_title,
                                                                               mv_last_measurements.unit            measured_unit,
                                                                               mv_last_measurements.value           last_measurement
from station
         inner join mv_first_measurements on station.id_station = mv_first_measurements.id_station
         join mv_last_measurements
              on mv_first_measurements.id_station = mv_last_measurements.id_station
where status = 'enabled'
order by station.id_station;
