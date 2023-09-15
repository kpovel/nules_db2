-- 1
with stationFirstMeasurment as (select id_station, min(time) as firstTimeMeasurement
                                from measurement
                                group by id_station)
select name, city
from station
         join stationFirstMeasurment sfm on station.id_station = sfm.id_station
where extract(year from sfm.firstTimeMeasurement) = 2022;

-- 2
select distinct city from station;

-- 3
select * from station where coordinates <-> point(35.058606,48.44803) = 0;

-- 4
with KyivsStations as (
    select id_station from station where city = 'Kyiv'
)

select measurement.*
from measurement
         join KyivsStations on KyivsStations.id_station = measurement.id_station
where date(time) between '2022-01-01' and '2022-02-20';

-- 5
select * from optimal_value;
