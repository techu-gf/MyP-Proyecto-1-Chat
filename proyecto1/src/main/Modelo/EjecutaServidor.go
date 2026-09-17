package main

import(
	"fmt"
	"flag"
	"net"
	"chat/src/main/Modelo/Servidor"
)

func main() {
	lecturaBandera := flag.Int("p", 1234, "Puerto para aceptar conexiones")

	flag.Parse()

	puerto := *lecturaBandera
	
	serv := servidor.CrearServidor(puerto)
	
	err := serv.Iniciar()

	if err != nil{
		fmt.Printf("Error al iniciar el servidor en el puerto %d", puerto)
		return
	}
}
