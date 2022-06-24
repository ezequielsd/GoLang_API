package main

// Projeto de uma api para persistir lista de nomes de ruas e a história da pessoa do nome.
// Usado o componentes:
// rota para API : https://github.com/gorilla/mux
// ORM para banco: https://gorm.io/index.html "go get gorm.io/driver/mysql" , para instalar driver do banco"
// Permitir acesso cors: https://github.com/gorilla/handlers "go get github.com/gorilla/handlers", para instalar

import (
	"api/database"
	"api/models"
	"api/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Dr. João Colin", Historia: "Nasceu em 2 de agosto de 1911, em Joinville/SC. Filho de Otto Colin e de Ingeborg Hermann Colin."},
		{Id: 2, Nome: "Emilio Landman", Historia: "Sei lá a estória"},
	}

	database.ConectaDataBase()
	fmt.Println("Iniciando servidor Rest em Go")
	routes.HandleRequest()
}


