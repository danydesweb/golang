package main

import (
	"net/http"

	"log"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "fmt"
    
    "strconv"
	"github.com/labstack/echo"

	
)


const token = "1061491150:AAHd2hlo9dPkxLajLWpsBzhCNc6XD_jg79w"
const mqttServer = "broker.shiftr.io"
const  mqttUser = "8a1ebb76"
const  mqttPassword = "ea34575201459ef"
const f = "Conectado a mqtt"

type Estado struct { 
	event_type  int
	dato int
} 

type Connected struct{
	
		id string
		tipo string
		 
		connection_id string
		connection_name string
		connection_read_only bool
		
	}


type User struct {
	gorm.Model `json:"model"`
	Name       string `json:"name"`
	Email      string `json:"email"`
}

type dbops struct {
	db *gorm.DB
}
func main() {
	//rutas

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	
	//start the server

	http.ListenAndServe(":3001", nil)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hola amigos"))
}



func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Contactos"))

	db, err := gorm.Open("postgres", "host=hansken.db.elephantsql.com port=5432 user=vtwlajng dbname=vtwlajng password=U27Dryhunm7bdKbAinQ5KcRXc894A-Nn")                                                                              
                                                                                                                                                                                                                                     
	if err != nil {                                                                                                                                                                                                                    
	  log.Print(err)
	  
	  initialMigration(db)
	  handleRequest(db)
	  handlerFunc("hola")
	 
	}                                                                                                                                                                                                                                  
																																																									   
	}
  
  
  func handlerFunc(msg string) func(echo.Context) error {
	  return func(c echo.Context) error {
		  return c.String(http.StatusOK, msg)
	  }
  }
  
  func allUsers(db *gorm.DB) func(echo.Context) error {
	  return func(c echo.Context) error {
		  var users []User
		  db.Find(&users)
		  fmt.Println("{}", users)
  
		  return c.JSON(http.StatusOK, users)
	  }
  }
  
  func newUser(db *gorm.DB) func(echo.Context) error {
	  return func(c echo.Context) error {
		  name := c.Param("name")
		  email := c.Param("email")
		  db.Create(&User{Name: name, Email: email})
		  return c.String(http.StatusOK, name+" user successfully created")
	  }
  }
  
  func deleteUser(db *gorm.DB) func(echo.Context) error {
	  return func(c echo.Context) error {
		  name := c.Param("name")
  
		  var user User
		  db.Where("name = ?", name).Find(&user)
		  db.Delete(&user)
  
		  return c.String(http.StatusOK, name+" user successfully deleted")
	  }
  }
  
  func updateUser(db *gorm.DB) func(echo.Context) error {
	  return func(c echo.Context) error {
		  name := c.Param("name")
		  email := c.Param("email")
		  var user User
		  db.Where("name=?", name).Find(&user)
		  user.Email = email
		  db.Save(&user)
		  return c.String(http.StatusOK, name+" user successfully updated")
	  }
  }
  
  func usersByPage(db *gorm.DB) func(echo.Context) error {
	  return func(c echo.Context) error {
		  limit, _ := strconv.Atoi(c.QueryParam("limit"))
		  page, _ := strconv.Atoi(c.QueryParam("page"))
		  var result []User
		  db.Limit(limit).Offset(limit * (page - 1)).Find(&result)
		  return c.JSON(http.StatusOK, result)
	  }
  }
  
  func handleRequest(db *gorm.DB) {
	  e := echo.New()
  
	  e.GET("/users", allUsers(db))
	  e.GET("/user", usersByPage(db))
	  e.POST("/user/:name/:email", newUser(db))
	  e.DELETE("/user/:name", deleteUser(db))
	  e.PUT("/user/:name/:email", updateUser(db))
  
	  e.Logger.Fatal(e.Start(":3001"))
  }
  
  func initialMigration(db *gorm.DB) {
  
	  db.AutoMigrate(&User{})
}
