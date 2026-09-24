package servidor

import (
	"testing"
)

//TestCreacion da un puerto específico al servidor
// y comprueba que se cree correctamente
func TestCreacion(t *testing.T) {
	puerto := 1234
	
	servidor := LevantarServidor(puerto)
	
	if servidor == nil {
		t.Errorf("No se instanció el objeto Servidor")
	}

	if servidor.GetPuerto() != puerto {
		t.Errorf("Esperado: %d, obtenido: %d", puerto, servidor.GetPuerto())
	}
}


