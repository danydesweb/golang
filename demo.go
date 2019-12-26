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

type struct temp{
	Grados int
	Centesimas int
}

type struct Dispositivo{
	nombre string
	tipo string
}

type struct event
func main(){
	event string
}




func main{
	
	func main{

	//crea a ClientOptions struct configuracion de la direccion broker, clientid, 
	opts := MQTT.NewClientOptions().AddBroker("tcp://mqtt.eclipse.org:1883")
	opts.SetClientID("go-simple1")
	opts.SetDefaultPublishHandler(f)
  
	//crea e incia  el cliente usando  ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
	  panic(token.Error())
	}
  
	//subscribe al topic /go-mqtt/sample and request messages to be delivered
   
	if token := c.Subscribe("go-mqtt/sample1", 0, nil); token.Wait() && token.Error() != nil {
	  fmt.Println(token.Error())
	  os.Exit(1)
	}
  
	//Publica 5 messages en /go-mqtt/sample 
  
	for i := 1; i < 5; i++ {
	  text := fmt.Sprintf("registro numero  #%d!", i)
	  token := c.Publish("go-mqtt/sample1", 0, false, text)
	  token.Wait()
	}
  
	time.Sleep(3 * time.Second)
  
	//borra suscripcion para  /go-mqtt/sample
	if token := c.Unsubscribe("go-mqtt/sample1"); token.Wait() && token.Error() != nil {
	  fmt.Println(token.Error())
	  os.Exit(1)
	}
  
	c.Disconnect(250)

}

}
