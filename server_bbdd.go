package main

import (
	"net/http"

	"log"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "fmt"
    
    "strconv"
	"github.com/labstack/echo"

	"flag"
	"time"
	"os"
	
	"encoding/json"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-telegram-bot-api/telegram-bot-api"
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
																																																									   
	log.Print("Conectado!")                                                                                                                                                                                                            
																																																									   
   // defer db.Close()    
  
   log.Printf("Hola!")
  
	  //
	  //mqtt
	  //
	  topic := flag.String("topic", "esp/test", "The topic name to/from which to publish/subscribe")
	  broker := flag.String("broker", "tcp://broker.shiftr.io:1883", "The broker URI. ex: tcp://10.10.1.1:1883")
	  password := flag.String("password", "try", "The password (optional)")
	  user := flag.String("user", "try", "The User (optional)")
	  id := flag.String("id", "fercho-command-center", "The ClientID (optional)")
	  cleansess := flag.Bool("clean", false, "Set Clean Session (default false)")
	  qos := flag.Int("qos", 0, "The Quality of Service 0,1,2 (default 0)")
	  // num := flag.Int("num", 1, "The number of messages to publish or subscribe (default 1)")
	   //payload := flag.String("message", "", "The message text to publish (default empty)")
	  // action := flag.String("action", "", "Action publish or subscribe (required)")
	  // store := flag.String("store", ":memory:", "The Store Directory (default use memory store)")
	  flag.Parse()
  
	  opts := MQTT.NewClientOptions().AddBroker("tcp://mqtt.eclipse.org:1883")
		opts.SetClientID("shiftr.io/Daniep")
		opts.SetDefaultPublishHandler( func(client MQTT.Client, msg MQTT.Message) {
		  fmt.Printf("TOPIC: %s\n", msg.Topic())
		  fmt.Printf("MSG: %s\n", msg.Payload())
		})
	  opts.AddBroker(*broker)
	  opts.SetClientID(*id)
	  opts.SetUsername(*user)
	  opts.SetPassword(*password)
	  opts.SetCleanSession(*cleansess)
	  
	  userEvents := make(chan string)
	  opts.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
		  userEvents <- string(msg.Payload())
  
		  // msg.Topic()
	  })
  
	  // simple mqtt pub
	  client := MQTT.NewClient(opts)
	  if token := client.Connect(); token.Wait() && token.Error() != nil {
		  panic(token.Error())
	  }
	  if token := client.Subscribe("esp/user", byte(*qos), nil); token.Wait() && token.Error() != nil {
		  log.Println(token.Error())
		  os.Exit(1)
		  c := MQTT.NewClient(opts)
		  if token := c.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		  }
		
		  //subscribe al topic /shiftr.io/Daniep/luces and request messages to be delivered
		 
		  if token := c.Subscribe("shiftr.io/Daniep/luces", 0, nil); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		  }
		
		  //Publica messages en /Daniep/luces
		
			
			text := fmt.Sprintf("las luces estan encendidas")
			token := c.Publish("shiftr.io/Daniep/luces", 0, false, text)
			token.Wait()
		  
		
		  time.Sleep(3600 * time.Second)
		
		  //borra suscripcion para  /Daniep/luces
		  if token := c.Unsubscribe("shiftr.io/Daniep/luces"); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		  }
		
		  c.Disconnect(250)
  
	  }
	  log.Println("Sample Publisher Started")
  
	  bot, err := tgbotapi.NewBotAPI(token)
	  if err != nil {
		  panic(err)
		  time.Sleep(3 * time.Second)
	  }
  
		  
	  
  
	  bot.Debug = true // Has the library display every request and response.
  
	  // Create a new UpdateConfig struct with an offset of 0.
	  // Future requests can pass a higher offset to ensure there aren't duplicates.
	  updateConfig := tgbotapi.NewUpdate(0)
  
	  // Tell Telegram we want to keep the connection open longer and wait for incoming updates.
	  // This reduces the number of requests that are made while improving response time.
	  updateConfig.Timeout = 60
  
	  // Now we can pass our UpdateConfig struct to the library to start getting updates.
	  // The GetUpdatesChan method is opinionated and as such, it is reasonable to implement
	  // your own version of it. It is easier to use if you have no special requirements though.
	  updates, err := bot.GetUpdatesChan(updateConfig)
  
	  // Now we're ready to start going through the updates we're given.
	  // Because we have a channel, we can range over it.
  
	  go func() {
		  for update := range updates {
			  // There are many types of updates. We only care about messages right now,
			  // so we should ignore any other kinds.
			  if update.Message == nil {
				  continue
			  }
  
			  // Sample #1
			  // // Because we have to create structs for every kind of request,
			  // // there's a number of helper functions to make creating common
			  // // types easier. Here, we're using the NewMessage helper which
			  // // returns a MessageConfig struct.
			  // msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
  
			  // // As there's too many fields for each Config to specify in a single
			  // // function call, we need to modify the result the helper gave us.
			  // msg.ReplyToMessageID = update.Message.MessageID
  
			  // // We're ready to send our message!
			  // // The Send method is for Configs that return a Message struct.
			  // // Sending Messages (among many other types) return a Message.
			  // // In this case, we don't care about the returned Message.
			  // // We only need to make sure our message went through successfully.
			  // if _, err := bot.Send(msg); err != nil {
			  // 	panic(err) // Again, this is a bad way to handle errors.
			  // }
  
			  // var byte payload = 1
			  // Sample #2
			  if update.Message.IsCommand() {
				  msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
				  switch update.Message.Command() {
				  case "calentar":
					  if update.Message.CommandArguments() != "" {
						  var payload []byte = []byte{0}
						  input := update.Message.CommandArguments()
						  temperatura, _ := strconv.Atoi(input)
						  payload[0] = byte(temperatura)
  
						  log.Printf("Calentar el agua a %v", payload)
						  token := client.Publish(*topic, byte(*qos), false, payload)
						  token.Wait()
						  msg.Text = "Ahora va..."
					  } else {
						  msg.Text = "A cuantos grados?"
					  }
				  case "help":
					  msg.Text = "type /sayhi or /status."
				  case "sayhi":
					  msg.Text = "Hi :)"
				  case "status":
					  msg.Text = "I'm ok."
  
				  case "encender":
					  msg.Text = "luces encendidas"					
						  encendido := Estado{event_type: 1, dato: 1,}
						  crear_json, _ := json.Marshal(encendido)
						  token := client.Publish(*topic, byte(*qos), false, crear_json)
						  token.Wait()
					  case "apagar":
						  msg.Text = "luces apagadas"					
							  apagado := Estado{event_type: 1, dato: 1,}
							  crear_json, _ := json.Marshal(apagado)
							  token := client.Publish(*topic, byte(*qos), false, crear_json)
							  token.Wait()
  
				  default:
					  msg.Text = "I don't know that command"
				  }
				  bot.Send(msg)
			  }
		  }
	  }()
  
	  for {
		  select {
		  case event := <-userEvents:
			  log.Println("Received User Event:")
			  log.Println(event)
  
			  chat_id, _ := strconv.Atoi(os.Getenv("TELEGRAM_CHAT_ID"))
			  msg := tgbotapi.NewMessage(int64(chat_id), event)
  
			  // // As there's too many fields for each Config to specify in a single
			  // // function call, we need to modify the result the helper gave us.
			  // msg.ReplyToMessageID = update.Message.MessageID
			  // TODO link to command?
  
			  // // We're ready to send our message!
			  // // The Send method is for Configs that return a Message struct.
			  // // Sending Messages (among many other types) return a Message.
			  // // In this case, we don't care about the returned Message.
			  // // We only need to make sure our message went through successfully.
			  if _, err := bot.Send(msg); err != nil {
				  panic(err) // Again, this is a bad way to handle errors.
			  }
		  }
		  time.Sleep(3 * time.Second)
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
  
	  e.Logger.Fatal(e.Start(":3000"))
  }
  
  func initialMigration(db *gorm.DB) {
  
	  db.AutoMigrate(&User{})
}
