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
	Users map[string]string `json:"users,omitempty"`
	Usernames []string `json:"usernames,omitempty"`
	Text string `json:"text,omitempty"`
	Roomname string `json:"roomname,omitempty"`
}

func CrearMensajeIdentify(username string) *Mensaje {
	return &Mensaje{
		Tipo : "IDENTIFY",
		Username : username,
	}
}

func CrearMensajeResponse(operation, result, extra string) *Mensaje {
	return &Mensaje{
		Tipo : "RESPONSE",
		Operation : operation,
		Result : result,
		Extra : extra,
	}
}

func CrearMensajeNewUser(username string) *Mensaje {
	return &Mensaje{
		Tipo : "NEW_USER",
		Username : username,
	}
}

func CrearMensajeStatus(status string) *Mensaje {
	return &Mensaje{
		Tipo : "STATUS",
		Status : status,
	}
}

func CrearMensajeNewStatus(username, status string) *Mensaje {
	return &Mensaje{
		Tipo : "NEW_STATUS",
		Username : username,
		Status : status,
	}
}

func CrearMensajeUsers() *Mensaje {
	return &Mensaje{
		Tipo : "USERS",
	}
}

func CrearMensajeUserList(users map[string]string) *Mensaje {
	return &Mensaje{
		Tipo : "USER_LIST",
		Users : users,
	}
}

func CrearMensajeText(username, text string) *Mensaje {
	return &Mensaje{
		Tipo : "TEXT",
		Username : username,
		Text : text,
	}
}

func CrearMensajeTextFrom(username, text string) *Mensaje {
	return &Mensaje{
		Tipo : "TEXT_FROM",
		Username : username,
		Text : text,
	}
}

func CrearMensajePublicText(text string) *Mensaje {
	return &Mensaje{
		Tipo : "PUBLIC_TEXT",
		Text : text,
	}
}

func CrearMensajePublicTextFrom(username, text string) *Mensaje {
	return &Mensaje{
		Tipo : "PUBLIC_TEXT_FROM",
		Username : username,
		Text : text,
	}
}

func CrearMensajeNewRoom(roomname string) *Mensaje {
	return &Mensaje{
		Tipo : "NEW_ROOM",
		Roomname : roomname,
	}
}

func CrearMensajeInvite(roomname string, usernames []string) *Mensaje {
	return &Mensaje{
		Tipo : "INVITE",
		Roomname : roomname,
		Usernames : usernames,
	}
}

func CrearMensajeInvitation(username, roomname string) *Mensaje {
	return &Mensaje{
		Tipo : "INVITATION",
		Username : username,
		Roomname : roomname,
	}
}

func CrearMensajeJoinRoom(roomname string) *Mensaje {
	return &Mensaje{
		Tipo : "JOIN_ROOM",
		Roomname : roomname,
	}
}

func CrearMensajeJoinedRoom(roomname, username string) *Mensaje {
	return &Mensaje{
		Tipo : "JOINED_ROOM",
		Roomname : roomname,
		Username : username,
	}
}

func CrearMensajeRoomUsers(roomname string) *Mensaje {
	return &Mensaje{
		Tipo : "ROOM_USERS",
		Roomname : roomname,
	}
}

func CrearMensajeRoomText(roomname, text string) *Mensaje {
	return &Mensaje{
		Tipo : "ROOM_TEXT",
		Roomname : roomname,
		Text : text,
	}
}

func CrearMensajeRoomTextFrom(roomname, username, text string) *Mensaje {
	return &Mensaje{
		Tipo : "ROOM_TEXT_FROM",
		Roomname : roomname,
		Username : username,
		Text : text,
	}
}

func CrearMensajeLeaveRoom(roomname string) *Mensaje {
	return &Mensaje{
		Tipo : "LEAVE_ROOM",
		Roomname : roomname,
	}
}

func CrearMensajeLeftRoom(roomname, username string) *Mensaje {
	return &Mensaje{
		Tipo : "LEFT_ROOM",
		Roomname : roomname,
		Username : username,
	}
}

func CrearMensajeDisconnect() *Mensaje {
	return &Mensaje{
		Tipo : "DISCONNECT",
	}
}

func CrearMensajeDisconnected(username string) *Mensaje {
	return &Mensaje{
		Tipo : "DISCONNECTED",
		Username : username,
	}
}

func (msg *Mensaje)EsValido() bool {
	return strings.TrimSpace(msg.Tipo) != ""
}

func (msg *Mensaje)GetTipo() string {
	return msg.Tipo
}

func (msg *Mensaje)GetUsername() string {
	return msg.Username
}

func (msg *Mensaje)GetOperation() string {
	return msg.Operation
}

func (msg *Mensaje)GetResult() string {
	return msg.Result
}

func (msg *Mensaje)GetExtra() string {
	return msg.Extra
}

func (msg *Mensaje)GetStatus() string {
	return msg.Status
}

func (msg *Mensaje)GetUsers() map[string]string {
	return msg.Users
}

func (msg *Mensaje)GetUsernames() []string {
	return msg.Usernames
}

func (msg *Mensaje)GetText() string {
	return msg.Text
}

func (msg *Mensaje)GetRoomname() string {
	return msg.Roomname
}
