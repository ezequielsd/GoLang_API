# Apresentação

Programa exemplo desenvolvido em Go Lang.
É uma API desenvolvida para persistir dados de cadastro de nomes de personalidades de ruas e sua breve história.


# Requisitos

Para executar o programa é necessário:
* Ter instalado o sdk do Go no seu sistema operacional. Pode ser pego em [https://go.dev/](https://go.dev/) 


# Recursos abordados

* Funções.
* Orientação de objetos.
* Modularização de packages.
* Utilização de rotas.
* MVC.
* Comunicação com banco de dados AWS em Mysql.
* CRUD completo em banco.
* Utilização de pacotes de terceiros.


# Pacotes adicionais

* Para rotas: https://github.com/gorilla/mux
* Para ORM para banco : https://gorm.io/index.html
* Driver Mysql para ORM: "go get gorm.io/driver/mysql"
* Configuração de acesso cors: https://github.com/gorilla/handlers


# Compilar e roda

Para compilar        : go build main.go
Para compilar e rodar: go run main.go

Para corrigir problemas de referencia de packages ou para possibilitar rodar packages fora do GO_HOME:

"go mod init"


# Script de banco

Na raiz tem um arquivo chamado "scriptDb.txt" para criação da base de dados e tabela.
