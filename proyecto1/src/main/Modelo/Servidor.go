package main

import(
	"fmt"
	//"flag"
	//"sync"
	"net"
)

func main() {
	puerto := 1234

	ln, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("No se activó la acción Listen debido al error:", err)
	}

	defer ln.Close()

	servidor := LevantarServidor(puerto)

	fmt.Println("Se levantó el servidor en el puerto", servidor.Puerto)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("No se pudo aceptar la conexión debido al error:", err)
		}

		go hiloCliente(conn)
	}
}
 
type Servidor struct{
	Puerto int
	Usuarios map[int]string
	Salas map[string]*string
}

func LevantarServidor(puertoDado int) *Servidor{
	return &Servidor{
		Puerto : puertoDado,
		Usuarios : make(map[int]string),
		Salas : make(map[string]*string),
	}
}

func hiloCliente(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Se conectó alguien:D")
}

func (serv *Servidor) recibeMensaje(){
	
}

func (serv *Servidor) enviaMensaje(){
	
}

func (serv *Servidor) identificaUsuario(nombre string) bool{
	return true
}

func (serv *Servidor) cambiaStatus(status string){
	
}

func (serv *Servidor) desconectaUsuario(nombre string){
	
}

func (serv *Servidor) crearSala(nombreSala, nombreUsuario string){
	
}

func (serv *Servidor) agregarUsuarioSala(nombreSala, nombreUsuario string){
	
}

func (serv *Servidor) verListaUsuarios(nombreSala string){
	
}

