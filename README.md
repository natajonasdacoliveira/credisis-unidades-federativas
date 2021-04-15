# credisis-unidades-federativas
Atividade 2 

# Instalando
git clone https://github.com/natajonasdacoliveira/credisis-unidades-federativas

cd /credisis-unidades-federativas

# Troubleshooting
Caso necessário, instalar os pacotes necessários com o comando go get.

Exemplo : go get github.com/labstack/echo/v4

# Banco de dados
Acesse o arquivo db/db.go e modifique de acordo com o banco utilizado.

"root:root@tcp(127.0.0.1:3306)/UF"

Em ordem: nome, senha, IP, nome do banco.

# API
Podem ser encontradas no arquivo server.go, na função main.

Exemplo de rota não autenticada:localhost:1323/municipios (GET)

Exemplo de rota autenticada:localhost:1323/jwt/municipios (POST, PUT, DELETE)

As rotas autenticadas utilizam o bearer token que é recebido na localhost:1323/login (POST)

Usuários:

{
	"email":"teste@gmail.com",
	"password":"123123Teste"
}

{
	"email":"teste2@gmail.com",
	"password":"123123Teste"
}

Utilize o body para enviar esses dados

# SQL

drop database if exists UF;
create database UF;
use uf;

create table user(	
	id int auto_increment primary key,
	email varchar(200) not null unique,
    password BINARY(60) not null
);

create table estados(
	id int auto_increment primary key,
	nome varchar(200) not null,
    sigla varchar(200) not null
);

create table municipios(
	id int auto_increment primary key,
	nome varchar(200) not null,
	prefeito varchar(200),
    populacao integer,
    id_estado_fk int,
    foreign key(id_estado_fk) references estados(id) on delete set null
);

insert into user values(null, "teste@gmail.com", "$2a$10$zHj6itdxGyYs9n3/nk6t6.UA6RyrsZIUe1waO.NQK6WaieFct0VXC");
insert into user values(null, "teste2@gmail.com", "$2a$10$zHj6itdxGyYs9n3/nk6t6.UA6RyrsZIUe1waO.NQK6WaieFct0VXC");

insert into estados values(null, "Rondônia", "RO");
insert into municipios values(null, "Ji-paraná", "Marcito Aparecido Pinto", 130.009, 1);










