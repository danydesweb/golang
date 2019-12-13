package main

import (
	"fmt"
	"log"
	"time"

	"github.com/yanzay/tbot"
)

const token = "1061491150:AAHd2hlo9dPkxLajLWpsBzhCNc6XD_jg79w"

func main() {
	//rutina := func(i int) {
	///fmt.Println("rutina 1", i, time.Now().Format("5"))

	time.Sleep(15 * time.Second)

	i := 0
	for i <= 10 {
		fmt.Println(i)
		i = i + 1

	}

	time.Sleep(10 * time.Second)

	bot := tbot.New(token)
	client := bot.Client()

	bot.HandleMessage("hola", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "Hola esto esta funcionando man !  ")
	})

	log.Fatal(bot.Start())

}
