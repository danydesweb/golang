package main

import(
	(
		"fmt"
		//import the Paho Go MQTT library
		MQTT "github.com/eclipse/paho.mqtt.golang"
		"os"
		"time"
	   
	  )
)

type struct DATOS{
	Grados int
	Centesimas int
	nombre string
	tipo string
	event string
	deviceType int
	eventType int


var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())




	
func main() {

		http.HandleFunc("/", )
		http.HandleFunc("/ambiente", "")
		http.HandleFunc("/sensores", "")
		http.ListenAndServe(":3005", nil)

	}
	
  
	c.Disconnect(250)

}
