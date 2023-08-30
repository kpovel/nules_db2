drop table if exists Favorite;
drop table if exists MQTT_Unit;
drop table if exists Measurement;
drop table if exists Station;
drop table if exists Optimal_Value;
drop table if exists Measured_Unit;
drop type if exists Units;
drop table if exists Category;
drop table if exists MQTT_Server;
drop type if exists Server_Status;

create type Server_Status as enum ('on', 'off');
create table MQTT_Server
(
    ID_Server char(10) primary key,
    Url       text,
    Status    Server_Status
);

create table Category
(
    ID_Category char(10) primary key,
    Designation varchar(255)
);

create type Units as ENUM ('nm', 'μm');
create table Measured_Unit
(
    ID_Measured_Unit char(10) primary key,
    Title            varchar(255),
    Unit             Units
);

create table Optimal_Value
(
    ID_Category      char(10),
    ID_Measured_Unit char(10),
    Bottom_Border    float,
    Upper_Border     float,

    foreign key (ID_Category) references Category (ID_Category),
    foreign key (ID_Measured_Unit) references Measured_Unit (ID_Measured_Unit)
);

create table Station
(
    ID_Station    char(10) primary key,
    City          varchar(50),
    Name          varchar(250),
    ID_SaveEcoBot char(10),
    ID_Server     char(10),
    Coordinates   point,

    foreign key (ID_Server) references MQTT_Server (ID_Server)
);

create table Measurement
(
    ID_Measurement   char(10),
    Time             timestamp,
    Value            decimal(8, 6),
    ID_Station       char(10),
    ID_Measured_Unit char(10),

    foreign key (ID_Station) references Station (ID_Station),
    foreign key (ID_Measured_Unit) references Measured_Unit (ID_Measured_Unit)
);

create table MQTT_Unit
(
    ID_Station       char(10),
    ID_Measured_Unit char(10),
    Message          varchar(255),
    "Order"          int,

    foreign key (ID_Station) references Station (ID_Station),
    foreign key (ID_Measured_Unit) references Measured_Unit (ID_Measured_Unit)
);

create table Favorite
(
    User_Name  varchar(255) primary key,
    ID_Station char(10),

    foreign key (ID_Station) references Station (ID_Station)
);
