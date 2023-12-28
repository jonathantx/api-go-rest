package main

import (
	"fmt"

	"github.com/jonathantx/go-rest-api/database"
	"github.com/jonathantx/go-rest-api/routes"
)

func main() {

	// models.Personalidades = []models.Personalidade{
	// 	{Id: 1, Nome: "Albert Einstein", Historia: "Albert Einstein foi um físico teórico alemão, que desenvolveu a teoria da relatividade geral, um dos pilares da física moderna ao lado da mecânica quântica."},
	// 	{Id: 2, Nome: "Santos Dummond", Historia: "Alberto Santos Dumont foi um aeronauta, esportista, autodidata e inventor brasileiro. Dumont é amplamente reconhecido como o pai da aviação, tendo realizado o primeiro voo homologado da história. Santos Dumont projetou, construiu e voou os primeiros balões dirigíveis com motor a gasolina."},
	// }

	database.ConnectionDataBase()

	fmt.Println("Iniciando Servidor!")
	routes.HandleRequest()

}
