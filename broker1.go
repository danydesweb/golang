package main

import (
  "fmt"
  //import the Paho Go MQTT library
  MQTT "github.com/eclipse/paho.mqtt.golang"
  "os"
  "time"
)

// define una función para el manejador de mensajes predeterminado
var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
  fmt.Printf("TOPIC: %s\n", msg.Topic())
  fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
  //crear una estructura ClientOptions configurando la dirección del agente, cliente

  opts := MQTT.NewClientOptions().AddBroker("tcp://mqtt.eclipse.org:1883")
  opts.SetClientID("go-sample")
  opts.SetDefaultPublishHandler(f)

  //crear e iniciar un cliente usando las ClientOptions 
  c := MQTT.NewClient(opts)
  if token := c.Connect(); token.Wait() && token.Error() != nil {
    panic(token.Error())
  }

  //subscribe al topic /go-mqtt/sample y solicita mensajes
  // espera el recibo para confirmar la suscripcion


  if token := c.Subscribe("go-mqtt/sample ", 0, nil); token.Wait() && token.Error() != nil {
    fmt.Println(token.Error())
    os.Exit(1)
  }

  //Publica  5 mensajes en /go-mqtt/sample 
  // el servidor después de envia cada mensaje
  for i := 0; i < 5; i++ {
    text := fmt.Sprintf("Yo soy Danielito,Fer me tomando sabes cuantas cervezas?????  #%d!", i)
    token := c.Publish("go-mqtt/sample", 0, false, text)
    token.Wait()
  }

  time.Sleep(3 * time.Second)

  //unsubscribe from /go-mqtt/sample
  if token := c.Unsubscribe("go-mqtt/sample"); token.Wait() && token.Error() != nil {
    fmt.Println(token.Error())
    os.Exit(1)
  }

  c.Disconnect(250)
}
