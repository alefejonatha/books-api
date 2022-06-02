# Books API
Esta API permite criar, editar, listar, pesquisar e deletar livros.

## Requisitos
* Go 1.18.1
* MySQL 8


## Dependências
Há várias dependências de terceiros usadas no projeto, entre elas as principais são o [Gin](https://github.com/gin-gonic/gin) e [Gorm](https://gorm.io/index.html). Abra o arquivo go.sum para obter detalhes das bibliotecas e versões utilizadas

## Construindo o projeto
No projeto foi utilizado o **MySQL** diretamente num container **Docker**, se você tiver um banco de dados MySQL configurado na sua máquina não esqueça de disponibilizar o serviço na porta 3306 ou alterar as configurações em **database/database.go**

Clone o projeto: 
```
https://github.com/alefejonatha/books-api.git
```

Acesse o diretório do projeto:
```
cd books-api 
```

Dentro do diretório do projeto digite:
```
go run main.go
```
## API endpoints

Método | URL | Retorno
:---: | --- | ---
GET | localhost:5000/api/v1/books | Lista com todos os livros
GET | localhost:5000/api/v1/books/{id} | Lista um livro por id
POST | localhost:5000/api/v1/books | Cadastra um novo livro
PUT | localhost:5000/api/v1/books | Substitui os dados de um livro
DELETE | localhost:5000/api/v1/books/{id} | Remove um livro
