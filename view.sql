create index idx_measurement_station_time on measurement (id_station, time);

create materialized view mv_first_measurements as
select distinct id_station,
                first_value(time) over (w_asc) as first_measured
from measurement
window w_asc as (partition by id_station order by time)
order by id_station;

create materialized view mv_last_measurements as
select distinct id_station,
                first_value(value) over (w_desc) as last_measurements
from measurement
window w_desc as (partition by id_station order by time desc)
order by id_station;
