package models

import "github.com/rafa1um/db"

type Carro struct {
	Id        int
	Titulo    string
	Marca     string
	Modelo    string
	Ano       int16
	Descricao string
	Preco     float32
}

func GetAllCars() []Carro {
	db := db.ConnectToDb()
	selectAllCars, err := db.Query("select * from carros")

	if err != nil {
		panic(err.Error())
	}

	carro := Carro{}
	carros := []Carro{}

	for selectAllCars.Next() {
		var id int
		var ano int16
		var titulo, marca, modelo, descricao string
		var preco float32

		err = selectAllCars.Scan(&id, &titulo, &marca, &modelo, &ano, &descricao, &preco)

		if err != nil {
			panic(err.Error())
		}
		carro.Titulo = titulo
		carro.Marca = marca
		carro.Modelo = modelo
		carro.Ano = ano
		carro.Descricao = descricao
		carro.Preco = preco

		carros = append(carros, carro)
	}
	defer db.Close()
	return carros
}
