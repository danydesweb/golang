package main

import (
  	"log"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "fmt"
    "net/http"
    "strconv"
    "github.com/labstack/echo"
)
                                                                                                                                                                                                                                  
/*type Users struct{
	
  id string
  name string
  mail string
  phone uint
  
} 

type Dispositivos struct{
	
  id string
  protocol string
  power_source string
  input_output string
  voltage uint
  
}*/

type User struct {
	gorm.Model `json:"model"`
	Name       string `json:"name"`
	Email      string `json:"email"`
}

type dbops struct {
	db *gorm.DB
}
func main() {                                                                                                                                                                                                                        
  db, err := gorm.Open("postgres", "host=hansken.db.elephantsql.com port=5432 user=vtwlajng dbname=vtwlajng password=U27Dryhunm7bdKbAinQ5KcRXc894A-Nn")                                                                              
                                                                                                                                                                                                                                     
  if err != nil {                                                                                                                                                                                                                    
    log.Print(err)
    
    initialMigration(db)
    handleRequest(db)
  }                                                                                                                                                                                                                                  
                                                                                                                                                                                                                                     
  log.Print("Conectado!")                                                                                                                                                                                                            
                                                                                                                                                                                                                                     
 // defer db.Close()                                                                                                                                                                                                                   
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

	e.Logger.Fatal(e.Start(":3000"))
}

func initialMigration(db *gorm.DB) {

	db.AutoMigrate(&User{})
}



