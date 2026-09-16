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
	
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", puerto))
	if err != nil {
		fmt.Println("No se activó la acción Listen debido al error:", err)
		return
	}
	
	defer ln.Close()
	
	serv := servidor.LevantarServidor(puerto)
	
	fmt.Println("Se levantó el servidor en el puerto", serv.GetPuerto())
	
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("No se pudo aceptar la conexión debido al error:", err)
			continue
		}

		go serv.ProcesoCliente(conn)
	}
}
