package main

import (
	"flag"
	"log"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
	"github.com/jinzhu/gorm"
 _ "github.com/jinzhu/gorm/dialects/postgres"
 "github.com/satori/go.uuid"
)

type Base struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
   }
   // BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) {
	uuid := uuid.NewV4()
	
	
	return scope.SetColumn("ID", uuid)
   }
   // User is the model for the user table.
   type User5 struct {
	Name   string    //`gorm:"column:name;size:128;not null;"`
	UserID uuid.UUID //`gorm:"type:uuid;not null"`
	Email   string    //`gorm:"column:email;size:128;not null;"`
	Profile  Profile
   }
   // Profile is the model for the profile table.
   type Profile struct {

	Name   string   // `gorm:"column:name;size:128;not null;"`
	UserID uuid.UUID //`gorm:"type:uuid;not null"`
   }
  
  
   var fetchedUser = &User5{}
	
func main() {

	log.Printf("Hola!")

	

	broker := flag.String("broker", "tcp://broker.shiftr.io:1883", "The broker URI. ex: tcp://10.10.1.1:1883")
	password := flag.String("password", "try", "The password (optional)")
	user := flag.String("user", "try", "The User (optional)")
	id := flag.String("id", "fercho-command-center", "The ClientID (optional)")
	cleansess := flag.Bool("clean", false, "Set Clean Session (default false)")
	
	flag.Parse()

	opts := MQTT.NewClientOptions()
	opts.AddBroker(*broker)
	opts.SetClientID(*id)
	
	opts.SetUsername(*user)
	opts.SetPassword(*password)
	opts.SetCleanSession(*cleansess)

	// simple mqtt pub
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())}

	db, err := gorm.Open("postgres", "host=hansken.db.elephantsql.com port=5432 user=vtwlajng dbname=vtwlajng password=	Sty59HjeuNLpFjjhRA5HNok1gHc58lVs")
	if err != nil {
	 panic(err)
	}
	db.LogMode(true)
	db.AutoMigrate(&User5{}, &Profile{})
	
	log.Println("Sample Publisher Started")

	bot, err := tgbotapi.NewBotAPI(("1061491150:AAHd2hlo9dPkxLajLWpsBzhCNc6XD_jg79w"))
	if err != nil {
		panic(err) // You should add better error handling than this!
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
	for update := range updates {
		// There are many types of updates. We only care about messages right now,
		// so we should ignore any other kinds.
		if update.Message == nil {
			continue
		}

	



		
		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
		

			case "signup":
				if update.Message.CommandArguments() != "" {
					
					input := update.Message.CommandArguments()
					
					
					user := &User5{Email: input}
						if db.Create(&user).Error != nil {
						log.Panic("Unable to create user.")
						}
					profile := &Profile{UserID: user.Base.ID}
					if db.Create(&profile).Error != nil {
					log.Panic("Unable to create profile.")
					}
					/*fetchedUser := &User5{}
					if db.Where("id = ?", profile.UserID).Preload("Profile").First(&fetchedUser).RecordNotFound() {
					log.Panic("Unable to find created user.")
					}
					fmt.Printf("User: %+v\n", fetchedUser)
					}*/
					log.Printf(input)

					

					msg.Text = "Usuario creado"
				} else {
					msg.Text = "Ingresa tu email"
					
						}

						/*func (base *Base) BeforeCreate(scope *gorm.Scope) error {
							uuid, err := uuid.NewV4()
							if err != nil {
							 return err
								}
							return scope.SetColumn("ID", uuid)
							   } 
						
						}*/
				bot.Send(msg)
				}
		
				msg.Text = "I don't know that command"
				
			}
		
		}

}