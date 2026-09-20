package servidor

import(
	"fmt"
	"net"
	"bufio"
	controlador "chat/src/main/Controlador/Servidor"
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
func (serv *Servidor)Iniciar() error {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", serv.puerto))
	if err != nil {
		return fmt.Errorf("No se pudo inciar el servidor debido al error: ", err)
	}
	
	defer ln.Close()

	fmt.Printf("Servidor activo en el puerto %d.\n", serv.puerto)
	
	go func() {
		for {
			conn, err := ln.Accept()

			if err != nil {
				return
			}
			
			go serv.ProcesoCliente(conn)
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
func (serv *Servidor)ProcesoCliente(conn net.Conn){
	defer conn.Close()
	
	fmt.Printf("Se conectó alguien desde la dirección %s\n", conn.RemoteAddr())
	
	scanner := bufio.NewScanner(conn)
	var usuario string
	ctrl := controlador.CrearControlador()
	
	for scanner.Scan(){
		mensaje := scanner.Bytes()

		msg, err := ctrl.MensajeSinJSON(mensaje)

		if err != nil{
			ctrl.OperacionInvalida(conn, "INVALID")
			continue
		}

		operacion, err := ctrl.ProcesaMensaje(msg, conn, usuario)

		if err != nil{
			continue
		}

		if operacion == "DISCONNECT"{
			if usuario != "" {
				serv.DesconectarUsuario(usuario)
				usuario = ""
			}
			return
		}else{
			errOp := serv.realizaOperacion(operacion, msg, conn)

			if operacion == "IDENTIFY" && errOp == nil{
				usuario = msg.GetUsername()
			} 
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error leyendo de %s: %v\n", conn.RemoteAddr(), err)
	}
	
	if usuario != "" {
		serv.DesconectarUsuario(usuario)
	}
}

func (serv *Servidor)realizaOperacion(operacion string, msg *mensaje.Mensaje, conn net.Conn) error{
	var respuesta *mensaje.Mensaje
	switch operacion{
		case "IDENTIFY":
		nombre := msg.GetUsername()

		err := serv.NuevoUsuario(nombre, conn)

		if err != nil{
			respuesta = mensaje.CrearMensajeResponse("IDENTIFY", "USER_ALREADY_EXISTS", nombre)
			serv.enviaMensaje(respuesta, conn)

			return err
		}else{
			respuesta = mensaje.CrearMensajeResponse("IDENTIFY", "SUCCESS", nombre)
			msgNuevoUsuario := mensaje.CrearMensajeNewUser(nombre)

			serv.enviaMensaje(respuesta, conn)
			serv.enviaMensaje(msgNuevoUsuario, conn)

			return nil
		}

		case "USERS":
		listaUsuario := serv.VerListaUsuarios()

		msgUserList := mensaje.CrearMensajeUserList(listaUsuario)

		serv.enviaMensaje(msgUserList, conn)
		
		return nil

		default:
		respuesta := mensaje.CrearMensajeResponse("INVALID", "INVALID", "")

		serv.enviaMensaje(respuesta, conn)
		
		return nil
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
}

func (serv *Servidor)Broadcast(conn net.Conn){
	
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

//La función enviaMensaje mandará un mensaje al cliente en caso de ser
//necesario.
func (serv *Servidor)enviaMensaje(mensaje *mensaje.Mensaje, conn net.Conn){
	ctrl := controlador.CrearControlador()
	
	bytesMensaje, _ := ctrl.MensajeAJSON(mensaje)

	ctrl.EnviarBytes(conn, bytesMensaje)
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

