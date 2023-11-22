-- Create a role with privileges
create role pavlo_admin with login password 'secure_password' createdb createrole;

-- Create a tablespace
create tablespace "eco-station-tablespace"
    owner pavlo_admin
    location '/Users/kpovel/Documents/nules/third_course/first_semester/db';

-- Create a new database template
create database "eco-station-template" owner pavlo_admin tablespace "eco-station-tablespace";

-- Create a template schema using ./schema.sql

-- Create a new_database with predefined schema
create database new_database
    owner pavlo_admin
    template "eco-station-template"
    tablespace "eco-station-tablespace";
