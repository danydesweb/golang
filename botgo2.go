package main

import (
	"log"
	//"time"
	//import the Paho Go MQTT library
   MQTT "github.com/eclipse/paho.mqtt.golang"
  "os"
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
		client.SendMessage(m.Chat.ID, "Hola soy SOL ! quieres acceder a alguna seccion de noticias? ")
	})
	bot.HandleMessage("si", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "Te ofrezco hoy deportes,politica,economia,futbol. Cual eliges")
	})
	bot.HandleMessage("no", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "te interesa conocer tu horoscopo ,el pronostico del tiempo o cotizacion del dolar?? Cual?")
	})
	
	bot.HandleMessage("deportes", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "https://argentina.as.com/")

	})

	bot.HandleMessage("dolar", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "http://www.dolarhoy.com/")

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

	bot.HandleMessage("/chau", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "Te espero pronto !")
	})
	bot.HandleMessage("/hola", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "Hola soy Daniel !")
	})
	bot.HandleMessage("/chau", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "Te espero pronto !")
	})
	bot.HandleMessage("horoscopo", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "https://www.univision.com/temas/horoscopo-diario")
	})
	bot.HandleMessage("/chau", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "Te espero pronto !")
	})
	bot.HandleMessage("/hola", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "Hola soy Daniel !")
	})

	bot.HandleMessage("broker", func(m *tbot.Message) {
		 broker
	})
	log.Fatal(bot.Start())

}


//define a function for the default message handler
var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
  }
  
  func broker() {
	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	opts := MQTT.NewClientOptions().AddBroker("tcp://mqtt.eclipse.org:1883")
	opts.SetClientID("go-simple1")
	opts.SetDefaultPublishHandler(f)
  
	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
	  panic(token.Error())
	}
  
	//subscribe to the topic /go-mqtt/sample and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	if token := c.Subscribe("go-mqtt/sample1", 0, nil); token.Wait() && token.Error() != nil {
	  fmt.Println(token.Error())
	  os.Exit(1)
	}
  
	//Publish 5 messages to /go-mqtt/sample at qos 1 and wait for the receipt
	//from the server after sending each message
	for i := 1; i < 5; i++ {
	  text := fmt.Sprintf("Fer Sabes cuantas cervezas ma tomare ahora??? #%d!", i)
	  token := c.Publish("go-mqtt/sample1", 0, false, text)
	  token.Wait()
	}
  
	time.Sleep(3 * time.Second)
  
	//unsubscribe from /go-mqtt/sample
	if token := c.Unsubscribe("go-mqtt/sample1"); token.Wait() && token.Error() != nil {
	  fmt.Println(token.Error())
	  os.Exit(1)
	}
  
	c.Disconnect(250)
  }




