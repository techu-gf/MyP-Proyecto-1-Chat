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

func CrearControlador(serv *servidor.Servidor) *Controlador{
	return &Controlador{
		serv: serv,
	}
}

//Crea un JSON con el mensaje dado
func (ctrl *Controlador)MensajeAJSON(msg *mensaje.Mensaje)([]byte, error){
	if msg == nil{
		return nil, fmt.Errorf("No se aceptan mensajes nulos")
	}

	return json.Marshal(msg)
}

//Interpreta el mensaje del JSON dado
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

		case "USERS":
		listaUsuarios := ctrl.serv.VerListaUsuarios()
		msgUserList := mensaje.CrearMensajeUserList(listaUsuarios)
		ctrl.EnviarMensaje(msgUserList, conn)

		return true

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

func (ctrl *Controlador)OperacionInvalida(conn net.Conn, resultado string){
	respuesta := mensaje.CrearMensajeResponse("INVALID", resultado, "")
	bytesRespuesta,_ := ctrl.MensajeAJSON(respuesta)
	ctrl.enviarBytes(conn, bytesRespuesta)
}

func (ctrl *Controlador)EnviarMensaje(msg *mensaje.Mensaje, conn net.Conn) error{
	bytesMensaje, err := ctrl.MensajeAJSON(msg)
	if err != nil{
		return fmt.Errorf("Error al crear el JSON.")
	}

	ctrl.enviarBytes(conn, bytesMensaje)
	return nil
}

func (ctrl *Controlador)enviarBytes(conn net.Conn, datos []byte){
	if !strings.HasSuffix(string(datos), "\n"){
		datos = append(datos, '\n')
	}
	conn.Write(datos)
}
