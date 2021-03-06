package main

import (
	"encoding/json"
	"net/http"
)

type Hijo struct {
	Edad   int
	Nombre string
}

func main() {
	http.HandleFunc("/", hijos)
	http.HandleFunc("/ambiente", ambientes)
	http.ListenAndServe(":3001", nil)
}

func hijos(w http.ResponseWriter, _ *http.Request) {
	ailen := Hijo{Edad: 19, Nombre: "ailen"}
	mateo := Hijo{Edad: 11, Nombre: "mateo"}
	json.NewEncoder(w).Encode(ailen)
	json.NewEncoder(w).Encode(mateo)
}

////////  segunda pagina ambientes 

type Ambiente struct {
	TempIdeal   int
	Nombre string
}


func ambientes(w http.ResponseWriter, _ *http.Request) {
	comedor := Ambiente{TempIdeal: 24, Nombre: "comedor"}
	cocina := Ambiente{TempIdeal: 22, Nombre: "cocina"}
	living := Ambiente{TempIdeal: 25, Nombre: "living"}
	habitacion1 := Ambiente{TempIdeal: 22, Nombre: "habitaciona"}
	habitacion2 := Ambiente{TempIdeal: 22, Nombre: "habitacion2"}
	habitacion3 := Ambiente{TempIdeal: 22, Nombre: "habitacion3"}
	baño := Ambiente{TempIdeal: 26, Nombre: "baño"}
	
	json.NewEncoder(w).Encode(comedor)
	json.NewEncoder(w).Encode(cocina)
	json.NewEncoder(w).Encode(living)
	json.NewEncoder(w).Encode(habitacion1)
	json.NewEncoder(w).Encode(habitacion2)
	json.NewEncoder(w).Encode(habitacion3)
	
	json.NewEncoder(w).Encode(baño)
}


///// sensores


type Sensores struct {
	Tipo   string
	Nombre string
	Ubicacion string
	Protocolo string
}

func sensores(w http.ResponseWriter, _ *http.Request) {
	sens_comedor := Sensores{Tipo: "temperatura", Nombre: "sens_comedor", Ubicacion: "sobre puerta",Protocolo: "mqtt"}
	sens_cocina := Sensores{Tipo: "temperatura", Nombre: "sens_cocina", Ubicacion: "sobre puerta",Protocolo: "mqtt"}
	sens_living :=Sensores{Tipo: "movimiento", Nombre: "sens_living", Ubicacion: "costado puerta",Protocolo: "mqtt"}
	sens_habitacion1 := Sensores{Tipo: "temperatura", Nombre: "sens_habitacion1", Ubicacion: "sobre puerta",Protocolo: "mqtt"}
	sens_habitacion2 := Sensores{Tipo: "temperatura", Nombre: "sens_habitacion2", Ubicacion: "sobre puerta",Protocolo: "mqtt"}
	sens_habitacion3 := Sensores{Tipo: "temperatura", Nombre: "sens_habitacion3", Ubicacion: "sobre puerta",Protocolo: "mqtt"}
	sens_baño := Sensores{Tipo: "temperatura/humedad", Nombre: "sens_baño", Ubicacion: "sobre puerta",Protocolo: "mqtt"}
	
	json.NewEncoder(w).Encode(sens_comedor)
	json.NewEncoder(w).Encode(sens_cocina)
	json.NewEncoder(w).Encode(sens_living)
	json.NewEncoder(w).Encode(sens_habitacion1)
	json.NewEncoder(w).Encode(sens_habitacion2)
	json.NewEncoder(w).Encode(sens_habitacion3)
	
	json.NewEncoder(w).Encode(sens_baño)
}
