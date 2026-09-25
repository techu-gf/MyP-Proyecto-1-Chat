package controlador

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"chat/src/main/Modelo/Mensaje"
	"chat/src/main/Modelo/Servidor"
)

//El control del lado del servidor nos ayudará a "interpretar" los
//mensajes que llegan en JSON recibidos a través del socket. Para esto
//se usará Mensaje.

type Controlador struct{
	serv *servidor.Servidor
}

//CrearControlador será nuestro "constructor" del Controlador.
func CrearControlador(serv *servidor.Servidor) *Controlador{
	return &Controlador{
		serv: serv,
	}
}

//Crea un JSON con el mensaje dado.
func (ctrl *Controlador)MensajeAJSON(msg *mensaje.Mensaje)([]byte, error){
	if msg == nil{
		return nil, fmt.Errorf("No se aceptan mensajes nulos")
	}

	return json.Marshal(msg)
}

//Interpreta el mensaje del JSON dado.
func (ctrl *Controlador)MensajeSinJSON(datosJson []byte)(*mensaje.Mensaje, error){
	var msg mensaje.Mensaje

	err := json.Unmarshal(datosJson, &msg)

	if err != nil{
		return nil, err
	}

	return &msg, nil
}

//Verifica que el mensaje tenga lo necesario y lo pasa al servidor para 
//que haga la acción solicitada. El booleano que regresa indica si el cliente
//podrá seguir activo o si se termina su conexión (se desconecta o comete algún
//error que haga que se termine).
func (ctrl *Controlador)ProcesaMensaje(msg []byte, conn net.Conn, usuario *string) bool{
	msgSinJSON, err := ctrl.MensajeSinJSON(msg)
	if err != nil || !msgSinJSON.EsValido(){
		ctrl.OperacionInvalida(conn, "INVALID")
		return false;
	}
	
	tipo := msgSinJSON.GetTipo()

	if *usuario == "" && tipo != "IDENTIFY"{
		ctrl.OperacionInvalida(conn, "NOT_IDENTIFIED")
		return false
	}

	switch tipo{
		case "IDENTIFY":
		nombre := msgSinJSON.GetUsername()
		
		if strings.TrimSpace(msgSinJSON.GetUsername()) == ""{
			ctrl.OperacionInvalida(conn, "INVALID")
			return false
		}
		
		if len(msgSinJSON.Username) > 8{
			ctrl.OperacionInvalida(conn, "INVALID")
			return false
		}

		err := ctrl.serv.NuevoUsuario(nombre, conn)

		if err != nil{
			respuesta := mensaje.CrearMensajeResponse("IDENTIFY", "USER_ALREADY_EXISTS", nombre)
			ctrl.EnviarMensaje(respuesta, conn)

			return false
		}else{
			*usuario = nombre
			respuesta := mensaje.CrearMensajeResponse("IDENTIFY", "SUCCESS", nombre)
			msgNuevoUsuario := mensaje.CrearMensajeNewUser(nombre)

			ctrl.EnviarMensaje(respuesta, conn)
			ctrl.serv.Broadcast(nombre, msgNuevoUsuario)
		}

		return true

		case "STATUS":
		nuevoStatus := msgSinJSON.GetStatus()

		if nuevoStatus == "ACTIVE" || nuevoStatus == "AWAY" || nuevoStatus == "BUSY"{
			ctrl.serv.CambiaStatus(*usuario, nuevoStatus)

			msgNewStatus := mensaje.CrearMensajeNewStatus(*usuario, nuevoStatus)
			ctrl.serv.Broadcast(*usuario, msgNewStatus)

			return true
		}

		ctrl.OperacionInvalida(conn, "INVALID")

		return true

		case "USERS":
		listaUsuarios := ctrl.serv.VerListaUsuarios()
		msgUserList := mensaje.CrearMensajeUserList(listaUsuarios)
		ctrl.EnviarMensaje(msgUserList, conn)

		return true

		case "TEXT":
		msgTextFrom := mensaje.CrearMensajeTextFrom(*usuario, msgSinJSON.Text)
		mandaMensaje := ctrl.serv.MensajeDirecto(msgSinJSON.Username, msgTextFrom)

		if !mandaMensaje{
			msgResponseText := mensaje.CrearMensajeResponse("TEXT", "NO_SUCH_USER", msgSinJSON.Username)
			ctrl.EnviarMensaje(msgResponseText, conn)
		}

		return true;

		case "PUBLIC_TEXT":
		msgPublicText := mensaje.CrearMensajePublicTextFrom(*usuario, msgSinJSON.GetText())
		ctrl.serv.Broadcast(*usuario, msgPublicText)

		return true;

		case "DISCONNECT":
		if *usuario != ""{
			ctrl.serv.DesconectarUsuario(*usuario)
			*usuario = ""
		}

		return false
		
		default:
		ctrl.OperacionInvalida(conn, "INVALID")
		return false
	}
}

//OperacionInvalida crea un mensaje del tipo Invalido, busca ahorrar el repetir
//este bloque de código en distintas operaciones.
func (ctrl *Controlador)OperacionInvalida(conn net.Conn, resultado string){
	respuesta := mensaje.CrearMensajeResponse("INVALID", resultado, "")
	bytesRespuesta,_ := ctrl.MensajeAJSON(respuesta)
	ctrl.enviarBytes(conn, bytesRespuesta)
}

//EnviarMensaje se encarga de transformar el mensaje dado a JSON y de mandar
//los bytes a la conexión dada. Busca evitar repetir el mismo bloque de código.
func (ctrl *Controlador)EnviarMensaje(msg *mensaje.Mensaje, conn net.Conn) error{
	bytesMensaje, err := ctrl.MensajeAJSON(msg)
	if err != nil{
		return fmt.Errorf("Error al crear el JSON.")
	}

	ctrl.enviarBytes(conn, bytesMensaje)
	return nil
}

//enviarBytes es un método privado que se encarga de únicamente mandar los JSON.
func (ctrl *Controlador)enviarBytes(conn net.Conn, datos []byte){
	if !strings.HasSuffix(string(datos), "\n"){
		datos = append(datos, '\n')
	}
	conn.Write(datos)
}
