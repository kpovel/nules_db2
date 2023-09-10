create type Server_Status as enum ('enabled', 'disabled');
create table MQTT_Server
(
    ID_Server serial primary key,
    Url       varchar(255) unique not null,
    Status    Server_Status       not null
);

create type Designation_Category as enum ('Excellent', 'Fine', 'Moderate', 'Poor', 'Very Poor', 'Severe');
create table Category
(
    ID_Category serial primary key,
    Designation Designation_Category not null
);

create type Units as ENUM ('%', 'mg/m3', 'hPa', 'Celsius', 'ppm', 'ppb', 'aqi', 'mkg/m3', 'n3v/god');
create table Measured_Unit
(
    ID_Measured_Unit serial primary key,
    Title            varchar(255) not null,
    Unit             Units        not null
);

create table Optimal_Value
(
    ID_Category      serial not null,
    ID_Measured_Unit serial not null,
    Bottom_Border    smallint,
    Upper_Border     smallint,

    foreign key (ID_Category) references Category (ID_Category),
    foreign key (ID_Measured_Unit) references Measured_Unit (ID_Measured_Unit)
);

create table Station
(
    ID_Station    serial primary key,
    City          varchar(50)     not null,
    Name          varchar(250)    not null,
    Status        Server_Status   not null,
    ID_SaveEcoBot char(32) unique not null,
    ID_Server     serial,
    Coordinates   point           not null,

    foreign key (ID_Server) references MQTT_Server (ID_Server)
);

create table Measurement
(
    ID_Measurement   serial primary key,
    Time             timestamp default now() not null,
    Value            decimal(12, 4)          not null,
    ID_Station       serial                  not null,
    ID_Measured_Unit serial                  not null,

    foreign key (ID_Station) references Station (ID_Station),
    foreign key (ID_Measured_Unit) references Measured_Unit (ID_Measured_Unit)
);

create table MQTT_Unit
(
    ID_Station       serial      not null,
    ID_Measured_Unit serial      not null,
    Message          varchar(32) not null,
    "Order"          serial,

    foreign key (ID_Station) references Station (ID_Station),
    foreign key (ID_Measured_Unit) references Measured_Unit (ID_Measured_Unit)
);

create table Favorite
(
    ID_User    serial,
    ID_Station serial not null,

    foreign key (ID_Station) references Station (ID_Station)
);
