package main()

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
  db, err := gorm.Open("vtwlajng", "host=
  	postgres://vbqlsvuc:qKDFKe1BvF1KlT8Itt7S8Xu5bAmoJkZA@rajje.db.elephantsql.com: port=5432 user=vtwlajng dbname=vtwlajng password=U27Dryhunm7bdKbAinQ5KcRXc894A-Nn")
  defer db.Close()
}



