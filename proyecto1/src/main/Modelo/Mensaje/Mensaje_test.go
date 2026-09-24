package mensaje

import "testing"

//TestTipo da un tipo de mensaje con varios espacios y comprueba
//que el programa quita los espacios innecesarios
func TestTipo(t *testing.T) {
	msg := &Mensaje{}

	msg.SetTipo("     IDENTIFY")
	obtenido := msg.GetTipo()
	if obtenido != "IDENTIFY"{
		t.Errorf("Type falló ya que se esperaba IDENTIFY y se obtuvo %s", obtenido)
	}
}

//TestUsername da un usuario con varios espacios y comprueba
//que el programa quita los espacios innecesariso
func TestUsername(t *testing.T) {
	msg := &Mensaje{}

	msg.SetUsername("Techu     ")
	obtenido := msg.GetUsername()
	if obtenido != "Techu"{
		t.Errorf("Username falló ya que se esperaba Techu y se obtuvo %s", obtenido)
	}
}

//TestTipo da un tipo de mensaje con varios espacios y comprueba
//que el programa quita los espacios innecesarios
func TestOperation(t *testing.T) {
	msg := &Mensaje{}

	msg.SetOperation("     IDENTIFY     ")
	obtenido := msg.GetOperation()
	if obtenido != "IDENTIFY"{
		t.Errorf("Operation falló ya que se esperaba IDENTIFY y se obtuvo %s", obtenido)
	}
}

//TestResult da un tipo de resultado con varios espacios y un salto
//de línea y comprueba que el programa quita los elementos innecesarios
func TestResult(t *testing.T) {
	msg := &Mensaje{}

	msg.SetResult("     \nSUCCESS    ")
	obtenido := msg.GetResult()
	if obtenido != "SUCCESS"{
		t.Errorf("Result falló ya que se esperaba SUCCESS y se obtuvo %s", obtenido)
	}
}

//TestRoomname da un nombre de cuarto con varios espacios y un salto
//de línea y comprueba que el programa quita los elementos innecesarios
func TestRoomname(t *testing.T) {
	msg := &Mensaje{}

	msg.SetRoomname("\n   \n  Sala 1\n")
	obtenido := msg.GetRoomname()
	if obtenido != "Sala 1"{
		t.Errorf("Roomname falló ya que se esperaba Sala 1 y se obtuvo %s", obtenido)
	}
}

//TestResult da un tipo de resultado con varios saltos de línea
//y comprueba que el programa quita los elementos innecesarios
func TestStatus(t *testing.T) {
	msg := &Mensaje{}

	msg.SetStatus("\nAWAY\n\n\n\n")
	obtenido := msg.GetStatus()
	if obtenido != "AWAY"{
		t.Errorf("Status falló ya que se esperaba AWAY y se obtuvo %s", obtenido)
	}
}
