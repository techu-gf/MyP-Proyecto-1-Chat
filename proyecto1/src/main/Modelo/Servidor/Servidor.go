package servidor

import(
	"fmt"
	"net"
	"bufio"
	"strings"
	"encoding/json"
	"chat/src/main/Modelo/Mensaje"
)

//Programa donde tendremos la estructura servidor. Esta estructura
//va a guardar el puerto en el cual estará escuchando, los usuarios
//que se conecten al servidor y las salas que posee. Se necargará
//de manejar los procesos de cada cliente a través de gorutines. 


//Estructura Servidor con el puerto donde se encuentra, el mapa
//de usuarios y el mapa de salas.
type Servidor struct{
	puerto int
	usuarios map[string]net.Conn
	status map[string]string
	salas map[string][]string
	acciones chan func() 
	broadcast chan []byte
}

//La función LevantarServidor da por iniciado el Servidor y se define 
//el puerto. Se crea el mapa vacío para usuarios y salas. 
func CrearServidor(puertoDado int) *Servidor{
	return &Servidor{
		puerto : puertoDado,
		usuarios : make(map[string]net.Conn),
		status : make(map[string]string),
		salas : make(map[string][]string),
		acciones : make(chan func()),
		broadcast : make(chan []byte),
	}
}

//La función Iniciar empieza la escucha continua en el puerto dado en
//busca de aceptar nuevos clientes a la vez que administra los procesos
//de nuevos usuarios, desconectar usuarios y recibo y envío de datos
func (serv *Servidor)Iniciar(procesarMensajeFunc func(msg []byte, conn net.Conn, usuario *string) bool) error {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", serv.puerto))
	if err != nil {
		return fmt.Errorf("No se pudo inciar el servidor debido al error: %v", err)
	}
	
	defer ln.Close()

	fmt.Printf("Servidor activo en el puerto %d.\n", serv.puerto)
	
	go func() {
		for {
			conn, err := ln.Accept()

			if err != nil {
				return
			}
			
			go serv.ProcesoCliente(conn, procesarMensajeFunc)
		}
	}()

	for{
		select{
			case accion := <- serv.acciones:
			accion()

			case datos := <- serv.broadcast:
			for _, conn := range serv.usuarios{
				conn.Write(datos)
			}
		}
	}
}

//La función ProcesoCliente se encargará de empezar el proceso personalizado
//del cliente. Va a recibir un mensaje del controlador, el mensaje vendrá
//por parte del cliente, procesará el mensaje dado y realizará la acción
//solicitada.
func (serv *Servidor)ProcesoCliente(conn net.Conn, procesarMensajeFunc func(msg []byte, conn net.Conn, usuario *string) bool){
	defer conn.Close()
	
	fmt.Printf("Se conectó alguien desde la dirección %s\n", conn.RemoteAddr())
	
	scanner := bufio.NewScanner(conn)
	var usuario string
	
	for scanner.Scan(){
		mensaje := scanner.Bytes()

		procesamientoMensaje := procesarMensajeFunc(mensaje, conn, &usuario)

		if !procesamientoMensaje{
			return
		}
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error leyendo de %s: %v\n", conn.RemoteAddr(), err)
	}
	
	if usuario != "" {
		serv.DesconectarUsuario(usuario)
	}
}

//La función NuevoUsuario asegurará que el username no está en uso y agregará
//al usuario al servidor junto con su estado predeterminado (ACTIVE).
func (serv *Servidor)NuevoUsuario(username string, conn net.Conn) error {
	respuesta := make(chan error)

	serv.acciones <- func(){
		if _,existe := serv.usuarios[username] ; existe{
			respuesta <- fmt.Errorf("El nombre %s ya está en uso.\n", username)
			return
		}

		serv.usuarios[username] = conn
		serv.status[username] = "ACTIVE"
		respuesta <- nil

		fmt.Printf("%s se ha conectado.\n", username)
	}

	return <- respuesta
}

//La función DesconectarUsuario buscará y eliminará del servidor al usuario
//que lo solicite.
func (serv *Servidor)DesconectarUsuario(username string){
	serv.acciones <- func(){
		if _,existe := serv.usuarios[username] ; existe{
			delete(serv.usuarios, username)
			delete(serv.status, username)

			fmt.Printf("%s se ha desconectado.\n", username)
		}
	}

	mensaje := mensaje.CrearMensajeDisconnected(username)
	serv.Broadcast(username, mensaje)
}

//La función Broadcast enviará mensajes a todos los clientes del servidor. Esta
//función se basa fuertemente en el proyecto https://github.com/Jayant-issar/go-tcp-chat.git
func (serv *Servidor)Broadcast(usuario string, msg *mensaje.Mensaje){
	serv.acciones <- func(){
		for _ , conn := range serv.usuarios{
			usuarioDestino := conn
			go func(c net.Conn){
				err := serv.enviaMensaje(msg, c)
				if err != nil{
					fmt.Printf("Error mandando mensaje al cliente: %v.\n", err)
				}
			}(usuarioDestino)
		}
	}
	
}

func (serv *Servidor)enviaMensaje(msg *mensaje.Mensaje, conn net.Conn) error{
	if msg == nil{
		return fmt.Errorf("No se puede mandar un mensaje nulo.\n")
	}

	bytesMensaje, err := json.Marshal(msg)

	if err != nil{
		return fmt.Errorf("Error al transformar el mensaje a JSON.\n")
	}

	if !strings.HasSuffix(string(bytesMensaje), "\n"){
		bytesMensaje = append(bytesMensaje, '\n')
	}

	_, err = conn.Write(bytesMensaje)

	if err != nil{
		return fmt.Errorf("Error al mandar mensaje por el socket.\n")
	}

	return nil
}

//La función GetPuertos regresa el puerto del servidor de forma que no podrá
//ser modificable.
func (serv *Servidor)GetPuerto() int{
	return serv.puerto
}

//La función GetUsuarios regresa el mapa de usuarios del servidor de forma que
//no podrá ser modificable.
func (serv *Servidor)GetUsuarios() map[string]net.Conn{
	return serv.usuarios
}

//La función GetSalas regresa el mapa de salas del servidor de forma que no
//podrá ser modificable.
func (serv *Servidor)GetSalas() map[string][]string{
	return serv.salas
}

//La función cambiaStatus cambiará el status mostrado del usuario que lo
//solicita.
func (serv *Servidor)CambiaStatus(status string){
	
}

//La función crearSala creará el cuarto que se solicita. 
func (serv *Servidor)CrearSala(nombreSala, nombreUsuario string){
	
}

//La función agregarUsuarioSala agregará al usuario solicitado a una sala
//especificada. 
func (serv *Servidor)AgregarUsuarioSala(nombreSala, nombreUsuario string){
	
}

//La función verListaUsuarios dará la lista de usuarios dentro de una sala.
func (serv *Servidor)VerListaUsuarios() map[string]string{
	respuesta := make(chan map[string]string)

	serv.acciones <- func(){
		lista := make(map[string]string)
		for nombre := range serv.usuarios{
			lista[nombre] = serv.status[nombre]
		}
		
		respuesta <- lista
	}

	return <- respuesta
}

