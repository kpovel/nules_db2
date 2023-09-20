-- 1
with stationFirstMeasurment as (select id_station, min(time) as firstTimeMeasurement
                                from measurement
                                group by id_station)
select name, city
from station
         join stationFirstMeasurment sfm on station.id_station = sfm.id_station
where extract(year from sfm.firstTimeMeasurement) >= 2021;

-- 2
select distinct city
from station;

-- 3
select *
from station
where coordinates <-> point(35.058606, 48.44803) = 0;

-- 4
with KyivsStations as (select id_station
                       from station
                       where city = 'Kyiv')

select measurement.id_measurement,
       measurement.id_station,
       measurement.id_measured_unit,
       measurement.time,
       measurement.value
from measurement
         join KyivsStations on KyivsStations.id_station = measurement.id_station
where date(time) between '2022-01-01' and '2022-02-20'
group by measurement.id_station, measurement.id_measured_unit, measurement.value, measurement.time,
         measurement.id_measurement;

-- 5
select measured_unit.title,
       category.designation,
       measured_unit.unit,
       optimal_value.bottom_border,
       optimal_value.upper_border
from optimal_value
         join category on category.id_category = optimal_value.id_category
         join measured_unit on measured_unit.id_measured_unit = optimal_value.id_measured_unit
group by measured_unit.title, category.designation, measured_unit.unit, optimal_value.bottom_border,
         optimal_value.upper_border;