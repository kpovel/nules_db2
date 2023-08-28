create table MQTT_Server
(
    ID_Server serial primary key,
    Url       text,
    Status    smallint
);

create table Category
(
    ID_Category serial primary key,
    Designation varchar(255)
);

create type Units as ENUM ('nm', 'μm');
create table Measured_Unit
(
    ID_Measured_Unit serial primary key,
    Title            varchar(255),
    Unit             Units
);

create table Optimal_Value
(
    ID_Category      serial,
    ID_Measured_Unit serial,

    foreign key (ID_Category) references Category (ID_Category),
    foreign key (ID_Measured_Unit) references Measured_Unit (ID_Measured_Unit)
);

create table Station
(
    ID_Station    serial check ( ID_Station > 0 ) primary key,
    City          varchar(50),
    Name          varchar(50),
    ID_SaveEcoBot serial,
    ID_Server     serial,

    foreign key (ID_Server) references MQTT_Server (ID_Server)
);

create table Measurement
(
    ID_Measurement   serial,
    Time             timestamp,
    Value            decimal(8, 6),
    ID_Station       serial,
    ID_Measured_Unit serial,

    foreign key (ID_Station) references Station (ID_Station),
    foreign key (ID_Measured_Unit) references Measured_Unit (ID_Measured_Unit)
);

create table MQTT_Unit
(
    ID_Station       serial,
    ID_Measured_Unit serial,
    Message          varchar(255),
    "Order"          int,

    foreign key (ID_Station) references Station (ID_Station),
    foreign key (ID_Measured_Unit) references Measured_Unit (ID_Measured_Unit)
);

create table Favorite
(
    User_Name  varchar(255) primary key,
    ID_Station serial,

    foreign key (ID_Station) references Station (ID_Station)
);

create table Coordinates
(
    ID_Station serial,
    Longitude  decimal(8, 6),
    Latitude   decimal(8, 6),

    foreign key (ID_Station) references Station (ID_Station)
);
