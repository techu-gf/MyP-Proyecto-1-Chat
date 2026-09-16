package controlador

import (
	"encoding/json"
	"fmt"
	"chat/src/main/Modelo/Mensaje"
	"chat/src/main/Modelo/Servidor"
)

//El control del lado del servidor nos ayudará a "interpretar" los
//mensajes que llegan en JSON recibidos a través del socket. Para esto
//se usará Mensaje.

type Controlador struct{
	serv *servidor.Servidor
}

func CrearControlador(servDado *servidor.Servidor) *Controlador{
	return &Controlador{
		serv : servDado,
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
//que haga la acción solicitada
func (ctrl *Controlador)ProcesaMensaje(msg *mensaje.Mensaje) error{
	if msg == nil || !msg.EsValido(){
		return fmt.Errorf("Mensaje inválido")
	}

	server := ctrl.serv

	switch msg.GetTipo(){
		case "IDENTIFY":
		server.IdentificaUsuario(msg.GetUsername())
		default:
		return fmt.Errorf("Tipo de mensaje inválido: %s", msg.GetTipo())
	}

	return nil
}
