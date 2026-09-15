package servidor

import(
	"fmt"
	"net"
	"bufio"
)
 
type Servidor struct{
	puerto int
	usuarios map[int]string
	salas map[string]*string
}

func LevantarServidor(puertoDado int) *Servidor{
	return &Servidor{
		puerto : puertoDado,
		usuarios : make(map[int]string),
		salas : make(map[string]*string),
	}
}

func (serv *Servidor)HiloCliente(conn net.Conn) {
	defer conn.Close()
	
	fmt.Printf("Se conectó alguien desde la dirección %s\n", conn.RemoteAddr())
	
	scanner := bufio.NewScanner(conn)
	
	for scanner.Scan(){
		mensaje := scanner.Text()

		if(mensaje) == "/salir"{
			fmt.Fprintln(conn, "Desconectando del servidor...")
			break
		}

		fmt.Printf("[%s] : %s\n", conn.RemoteAddr(), mensaje)
	}

	
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error leyendo de %s: %v\n", conn.RemoteAddr(), err)
	}

	fmt.Printf("Cliente %s desconectado.\n", conn.RemoteAddr())
}

func (serv *Servidor) GetPuerto() int{
	return serv.puerto
}

func (serv *Servidor) GetUsuarios() map[int]string{
	return serv.usuarios
}

func (serv *Servidor) GetSalas() map[string]*string{
	return serv.salas
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

