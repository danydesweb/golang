package main

import (
	"log"
  "time"

	"github.com/yanzay/tbot"
)

const token = "1061491150:AAHd2hlo9dPkxLajLWpsBzhCNc6XD_jg79w"

func main() {
	bot := tbot.New(token)
	client := bot.Client()

	bot.HandleMessage("/hola", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "Hola soy Daniel !")
	})
  bot.HandleMessage("/chau", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "Te espero pronto !")
	})

	log.Fatal(bot.Start())


}

func alarm() {
  bot := tbot.New(token)
  client := bot.Client()
  hora := time.Now()

    if hora == Hour(15:20:00,){

      client.SendMessage(m.Chat.ID, "Alarma")}


}
