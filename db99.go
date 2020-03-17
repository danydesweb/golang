package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Admin struct {
	gorm.Model
	Name string
	Age  uint  //`gorm:"default:18"`
  }
  
  
  
func main() {
	db, err := gorm.Open("postgres", "host=hansken.db.elephantsql.com port=5432 user=vtwlajng dbname=vtwlajng password=	Sty59HjeuNLpFjjhRA5HNok1gHc58lVs")

	if err != nil {
		log.Print(err)


	

	}

	var juan = Admin{Name:"juan" , Age: 45}
	db.Create(&juan)

	log.Print("Conectado!")
	
	defer db.Close()
}