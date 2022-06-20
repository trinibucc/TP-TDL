create database Espectadores;
use Espectadores;

drop table if exists Compradores;

create table Compradores(
Id int unsigned auto_increment primary key not null,
Nombre varchar(255) not null,
Compra double
);