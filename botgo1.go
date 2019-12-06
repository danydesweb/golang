package main

import (
	//"log"
  "time"
  "fmt"


	"github.com/yanzay/tbot"
  //"github.com/lestrrat-go/strftime"
)

const token = "1061491150:AAHd2hlo9dPkxLajLWpsBzhCNc6XD_jg79w"

func main() {
	bot := tbot.New(token)
	client := bot.Client()

	bot.HandleMessage("hola", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "Hola soy Daniel !  ")
	})
  bot.HandleMessage("bien", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "yo tambien ! que haces?")
	})
  bot.HandleMessage("mal", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "Por que ? que ocurre?")
	})
  bot.HandleMessage("mirando tele", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "que miras?")
	})
  bot.HandleMessage("deportes", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "https://argentina.as.com/")
	})
  bot.HandleMessage("politica", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "https://www.politicargentina.com/")
	})
  bot.HandleMessage("economia", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "https://www.infobae.com/economia/")
	})
  bot.HandleMessage("futbol", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "https://www.futbolparatodos.com.ar/primera-division-superliga")
	})
  bot.HandleMessage("/hola", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "Hola soy Daniel !")
	})
  bot.HandleMessage("/chau", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "Te espero pronto !")
	})
  bot.HandleMessage("/hola", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "Hola soy Daniel !")
	})
  bot.HandleMessage("/chau", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "Te espero pronto !")
	})
  bot.HandleMessage("/hola", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "Hola soy Daniel !")
	})
  bot.HandleMessage("/chau", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "Te espero pronto !")
	})
  bot.HandleMessage("/hola", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "Hola soy Daniel !")
	})




	log.Fatal(bot.Start())


}

func alarma(){
  t:= time.Now()
  if  t.Format(Hour > 12):{
      client.SendMessage(m.Chat.ID, "Alarma")}


}
/* func main() {

p := fmt.Println
t:= time.Now()
p(t.Format(time.RFC3339))

} */
