package mensaje

import "strings"

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

func esValido(msg *Mensaje) bool{
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
