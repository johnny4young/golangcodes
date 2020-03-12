package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Producto struct {
	gorm.Model
	CodigoBarras string
	Precio       uint
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=root password=password dbname=dbtest sslmode=disable")
	if err != nil {
		panic("Error en la conexiòn a la BD " + err.Error())
	}
	defer db.Close()

	fmt.Println("se conectó a la base de datos")

	// db.CreateTable(&Producto{})

	// db.Create(&Producto{
	// 	CodigoBarras: "11011101101",
	// 	Precio:       1500,
	// })

	var p Producto
	db.First(&p)
	fmt.Println(p)
}
