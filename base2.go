package main

import (
	"log"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Dispositivo struct {
  gorm.Model
  Dispositivo string
  Codigo uint
  Tipo  string
  Ubicacion string
}

func main() {

	log.Printf("Conectando a base de datos!")

  db, err := gorm.Open("vtwlajng", "U27Dryhunm7bdKbAinQ5KcRXc894A-Nn")
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()

  // Migrate the schema
  db.AutoMigrate(&Dispositivo{})

  // Create
  db.Create(&Dispositivo{Dispositivo: "sensor luz", Codigo: 1000,Tipo: "Input", Ubicacion: "Patio"})

  // Read
  var Dispositivo Dispositivo1
  db.First(&product, 1) // find product with id 1
  db.First(&Dispositivo, Dispositivo: "sensor luz", Codigo: 1000,Tipo: "Input", Ubicacion: "Patio") // find product with code l1212

  // Update - update product's price to 2000
  db.Model(&product).Update("Price", 2000)

  // Delete - delete product
  db.Delete(&product)
}
