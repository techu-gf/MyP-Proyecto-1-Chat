package main

import(
	"fmt"
	"flag"
	"chat/src/main/Modelo/Servidor"
	controlador "chat/src/main/Controlador/Servidor"
)

func main() {
	lecturaBandera := flag.Int("p", 1234, "Puerto para aceptar conexiones")

	flag.Parse()

	puerto := *lecturaBandera
	
	serv := servidor.CrearServidor(puerto)
	
	ctrl := controlador.CrearControlador(serv)

	err := serv.Iniciar(ctrl.ProcesaMensaje)

	if err != nil{
		fmt.Printf("Error al iniciar el servidor en el puerto %d", puerto)
		return
	}
}
