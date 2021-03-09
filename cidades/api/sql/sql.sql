CREATE DATABASE IF NOT EXISTS pesquisas;
USE pesquisas;

DROP TABLE IF EXISTS cidades;

CREATE TABLE cidades(
    id int auto_increment primary key,
	nome varchar (60) not null,	
	cpf varchar (11) not null,
	telefone varchar (20) not null,
	email varchar (50) not null,
	rua varchar (50) not null,
	complemento varchar (50) not null,
	codigo varchar (5) not null
	);
