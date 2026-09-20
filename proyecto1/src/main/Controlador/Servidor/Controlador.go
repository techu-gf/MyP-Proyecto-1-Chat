package controlador

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"chat/src/main/Modelo/Mensaje"
)

//El control del lado del servidor nos ayudará a "interpretar" los
//mensajes que llegan en JSON recibidos a través del socket. Para esto
//se usará Mensaje.

type Controlador struct{}

func CrearControlador() *Controlador{
	return &Controlador{}
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
//que haga la acción solicitada
func (ctrl *Controlador)ProcesaMensaje(msg *mensaje.Mensaje, conn net.Conn, usuario string)(string, error){
	if msg == nil || !msg.EsValido(){
		ctrl.OperacionInvalida(conn, "INVALID")
		return "", fmt.Errorf("Mensaje inválido")
	}

	tipo := msg.GetTipo()

	if usuario == "" && tipo != "IDENTIFY"{
		ctrl.OperacionInvalida(conn, "NOT_IDENTIFIED")
		return "", fmt.Errorf("El usuario intentó actuar sin haberse identificado.")
	}

	switch tipo{
		case "IDENTIFY":
		username := msg.GetUsername()

		if strings.TrimSpace(username) == ""{
			ctrl.OperacionInvalida(conn, "INVALID")
			return "", fmt.Errorf("Usuario nulo.")
		}

		return "IDENTIFY", nil

		case "USERS":
		if usuario == ""{
			ctrl.OperacionInvalida(conn, "INVALID")
			return "", fmt.Errorf("Usuario nulo.")
		}

		return "USERS", nil

		case "PUBLIC_TEXT":
		if usuario == ""{
			ctrl.OperacionInvalida(conn, "INVALID")
			return "", fmt.Errorf("Usuario nulo.")
		}

		return "PUBLIC_TEXT", nil
		

		case "DISCONNECT":
		return "DISCONNECT", nil
		
		default:
		ctrl.OperacionInvalida(conn, "INVALID")
		return "", fmt.Errorf("Tipo de mensaje inválido.")
	}
}

func (ctrl *Controlador)OperacionInvalida(conn net.Conn, resultado string){
	respuesta := mensaje.CrearMensajeResponse("INVALID", resultado, "")
	bytesRespuesta,_ := ctrl.MensajeAJSON(respuesta)
	ctrl.EnviarBytes(conn, bytesRespuesta)
}

func (ctrl *Controlador)EnviarBytes(conn net.Conn, datos []byte){
	if !strings.HasSuffix(string(datos), "\n"){
		datos = append(datos, '\n')
	}
	conn.Write(datos)
}
