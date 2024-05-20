# Projeto Eulabs de produtos #

Este projeto teste fornce APIs para listar, encontrar, atualizar e apagar produtos em tabela de banco de dados.

O projeto foi construído com a linguagem [Go](https://go.dev/), utilizando o framework [Echo](https://echo.labstack.com/) com os dados armazenados num banco de dados [Mysql](https://www.mysql.com/)

Para executar a aplicação deve-se primeiro executar o comand [docker compose](docker-compose.yml) na raiz do projeto. Esse comando subirá um container Docker com um banco de dados Mysql. O script de inicialização do banco de dados e das tabelas encontra-se na pasta [init_database](init_database/all.sql)

Após o banco de dados ser inicializado, deve-se executar o comando [make](makefile) na raiz do projeto.

A aplicação subirá com a porta designada no arquivo [config](config.yaml).

Segue também os esquemas do [Postman](postman/Eulabs.postman_collection.json) que deverão ser importados e utilizados para executar os endopints da aplicação.

A aplicação precisa de um token. Esse token deverá ser passado para cada endpoint chamado.
 


