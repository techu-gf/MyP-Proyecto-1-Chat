package servidor

import(
	"fmt"
	"net"
	"bufio"
)

//Programa donde tendremos la estructura servidor. Esta estructura
//va a guardar el puerto en el cual estará escuchando, los usuarios
//que se conecten al servidor y las salas que posee. Se necargará
//de manejar los procesos de cada cliente a través de gorutines. 


//Estructura Servidor con el puerto donde se encuentra, el mapa
//de usuarios y el mapa de salas.
type Servidor struct{
	puerto int
	usuarios map[int]net.Conn
	salas map[string][]string
}

//La función LevantarServidor da por iniciado el Servidor y se define 
//el puerto. Se crea el mapa vacío para usuarios y salas. 
func LevantarServidor(puertoDado int) *Servidor{
	return &Servidor{
		puerto : puertoDado,
		usuarios : make(map[int]net.Conn),
		salas : make(map[string][]string),
	}
}

//La función ProcesoCliente se encargará de empezar el proceso personalizado
//del cliente. Va a recibir un mensaje del controlador, el mensaje vendrá
//por parte del cliente, procesará el mensaje dado y realizará la acción
//solicitada.
func (serv *Servidor)ProcesoCliente(conn net.Conn) {
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

//La función GetPuertos regresa el puerto del servidor de forma que no podrá
//ser modificable.
func (serv *Servidor) GetPuerto() int{
	return serv.puerto
}

//La función GetUsuarios regresa el mapa de usuarios del servidor de forma que
//no podrá ser modificable.
func (serv *Servidor) GetUsuarios() map[int]net.Conn{
	return serv.usuarios
}

//La función GetSalas regresa el mapa de salas del servidor de forma que no
//podrá ser modificable.
func (serv *Servidor) GetSalas() map[string][]string{
	return serv.salas
}

//La función enviaMensaje mandará un mensaje al cliente en caso de ser
//necesario.
func (serv *Servidor) enviaMensaje(){
	
}

//La función NuevoUsuario recibe la solicitud de agregar un usuario al
//servidor.
func (serv *Servidor) NuevoUsuario(nombreUsuario string, conexion net.Conn){
	
}

//La función identificaUsuario asegura que el nombre del usuario es válido
//y que el usuario se podrá unir al servidor.
func (serv *Servidor) IdentificaUsuario(nombre string){
	fmt.Printf("Se quiso identificar como %s", nombre)
}

//La función cambiaStatus cambiará el status mostrado del usuario que lo
//solicita.
func (serv *Servidor) CambiaStatus(status string){
	
}

//La función desconectaUsuario va a desconectar al usuario del servidor.
func (serv *Servidor) DesconectaUsuario(nombre string){
	
}

//La función crearSala creará el cuarto que se solicita. 
func (serv *Servidor) CrearSala(nombreSala, nombreUsuario string){
	
}

//La función agregarUsuarioSala agregará al usuario solicitado a una sala
//especificada. 
func (serv *Servidor) AgregarUsuarioSala(nombreSala, nombreUsuario string){
	
}

//La función verListaUsuarios dará la lista de usuarios dentro de una sala.
func (serv *Servidor) VerListaUsuarios(nombreSala string){
	
}

