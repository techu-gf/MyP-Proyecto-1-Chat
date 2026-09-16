package mensaje

import "strings"

//Mensaje será el programa del lado del servidor que nos ayude a
//manejar los mensajes. Guardará la información necesaria del
//mensaje y nos permitirá acceder a esta

//Mensaje con los campos posibles que puede tener el mensaje JSON.
type Mensaje struct{
	Tipo string `json:"type,omitempty"`
	Username string `json:"username,omitempty"`
	Operation string `json:"operation,omitempty"`
	Result string `json:"result,omitempty"`
	Extra string `json:"extra,omitempty"`
	Status string `json:"status,omitempty"`
	Users string `json:"users,omitempty"`
	Text string `json:"text,omitempty"`
	Roomname string `json:"roomname,omitempty"`
}

func CrearMensajeIdentify(username string) *Mensaje{
	return &Mensaje{
		Tipo : "IDENTIFY",
		Username : username,
	}
}

func CrearMensajeResponse(operation, result, extra string) *Mensaje{
	return &Mensaje{
		Tipo : "RESPONSE",
		Result : result,
		Extra : extra,
		
	}
}

func CrearMensajeNewUser(username string) *Mensaje{
	return &Mensaje{
		Tipo : "NEW_USER",
		Username : username,
	}
}

func CrearMensajeStatus(status string) *Mensaje{
	return &Mensaje{
		Tipo : "STATUS",
		Status : status,
	}
}

func CrearMensajeNewStatus(status, username string) *Mensaje{
	return &Mensaje{
		Tipo : "NEW_STATUS",
		Username : username,
		Status : status,
	}
}

func CrearMensajeUsers() *Mensaje{
	return &Mensaje{
		Tipo : "USERS",
	}
}

func CrearMensajeUSER_LIST(username string) *Mensaje{
	return &Mensaje{
		Tipo : "IDENTIFY",
		Username : username,
	}
}

func (msg *Mensaje) EsValido() bool{
	return strings.TrimSpace(msg.Tipo) != ""
}

func (msg *Mensaje) GetTipo() string {
	return msg.Tipo
}

func (msg *Mensaje) GetUsername() string {
	return msg.Username
}

func (msg *Mensaje) GetOperation() string {
	return msg.Operation
}

func (msg *Mensaje) GetResult() string {
	return msg.Result
}

func (msg *Mensaje) GetExtra() string {
	return msg.Extra
}

func (msg *Mensaje) GetStatus() string {
	return msg.Status
}

func (msg *Mensaje) GetUsers() string {
	return msg.Users
}

func (msg *Mensaje) GetText() string {
	return msg.Text
}

func (msg *Mensaje) GetRoomname() string {
	return msg.Roomname
}

func (msg *Mensaje) SetTipo(tipo string) {
	msg.Tipo = strings.TrimSpace(tipo)
}

func (msg *Mensaje) SetUsername(username string) {
	msg.Username = strings.TrimSpace(username)
}

func (msg *Mensaje) SetOperation(operation string) {
	msg.Operation = strings.TrimSpace(operation)
}

func (msg *Mensaje) SetResult(result string) {
	msg.Result = strings.TrimSpace(result)
}

func (msg *Mensaje) SetExtra(extra string) {
	msg.Extra = extra
}

func (msg *Mensaje) SetStatus(status string) {
	msg.Status = strings.TrimSpace(status)
}

func (msg *Mensaje) SetUsers(users string) {
	msg.Users = users
}

func (msg *Mensaje) SetText(text string) {
	msg.Text = text
}

func (msg *Mensaje) SetRoomname(roomname string) {
	msg.Roomname = strings.TrimSpace(roomname)
}
