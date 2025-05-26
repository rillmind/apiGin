# apiGin

Esse projeto é uma api onde eu usarei para colocar em prática os meus conhecimentos que eu aprenderei estudando Go.

## Pré-requisitos

- [Go](https://golang.org/doc/install) 1.23.3
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/)

## Instalação

clone o projeto em sua máquina:

```sh
git clone https://github.com/rillmind/apiGin
cd apiGin
```

com go instalado:

```sh
go install
go run ./cmd
```

com docker:

```sh
docker build -t goapi .
docker compose up -d
```

Confira se se o host condiz com como você deseja rodar a aplicação.
se for rodar com docker compose use 'host = docker', mas se for rodar
a api localmente use 'host = local' em [conn.go](./db/conn.go).

## Rotas

Você pode conferir as rotas para teste no diretorio router:

- [Rotas de 'user': ](./router/user.go)
- [Rotas de 'product': ](./router/product.go)
