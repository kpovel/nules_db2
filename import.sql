\COPY MQTT_Server (Url, Status, ID_Server) FROM 'data/input/MQTT_Server.csv' WITH (FORMAT csv, DELIMITER ',', HEADER true);
\COPY Category FROM 'data/input/Category.csv' WITH (FORMAT csv, DELIMITER ',', HEADER true);
\copy measured_unit (Title, Unit, ID_Measured_Unit) from 'data/input/Measured_Unit.csv' with (format csv, delimiter ',', header true);
\copy optimal_value (Bottom_Border, Upper_Border, ID_Category, ID_Measured_Unit) from 'data/output/optimal_value.csv' with (format csv, delimiter ',', header true);
\copy station (ID_Station, City, Name, Status, ID_SaveEcoBot, ID_Server, Coordinates) from 'data/output/station.csv' with (format csv, delimiter ',', header true);
\copy measurement from 'data/output/results.csv' with (format csv, delimiter ',', header true);
\copy mqtt_unit from 'data/output/mqtt_unit.csv' with (format csv, delimiter ',', header true);
\copy favorite from 'data/input/Favorite_Station.csv' with (format csv, delimiter ',', header true);