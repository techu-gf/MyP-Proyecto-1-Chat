package servidor

import(
	"net"
)

//Programa donde tendremos la estructura cuarto. Cada cuarto tendrá
//su lista de usuarios y el nombre del cuarto.

//Estructura Cuarto con el nombre del cuarto y su lista de usuarios
type Cuarto struct{
	nombre string
	usuarios map[string]Cliente
	listaInvitados map[string]Cliente
}

//La función CrearCuarto inicializa los atributos del cuarto. 
func CrearCuarto(nombreCuarto string) *Cuarto{
	return &Cuarto{
		nombre : nombreCuarto,
		usuarios : make(map[string]Cliente),
		listaInvitados : make(map[string]Cliente),
	}
}

//La función AgregarUsuarioSala agrega al usuario a la sala después de
//que haya aceptado la invitación. Regresa true si el usuario estaba en
//la lista de invitados, false en otros casos.
func (cuarto *Cuarto)AgregarUsuarioSala(nombreUsuario, status string, conn net.Conn) bool{
	if len(cuarto.usuarios) == 0{
		cuarto.usuarios[nombreUsuario] = Cliente{
			conn: conn,
			status: status,
		}
		return true
	}

	if _, existe := cuarto.listaInvitados[nombreUsuario]; existe{
		cuarto.usuarios[nombreUsuario] = Cliente{
			conn: conn,
			status: status,
		}
		return true
	}

	return false

}

func (cuarto *Cuarto)AgregarInvitado(nombreUsuario, status string, conn net.Conn){
	cuarto.listaInvitados[nombreUsuario] = Cliente{
		conn: conn,
		status: status,
	}
}

//La función EliminarUsuario elimina al usuario de la sala y regresa
//true si la sala está vacía (no hay más usuarios) o false si todavía
//quedan usuarios en ella. 
func (cuarto *Cuarto)EliminarUsuario(nombreUsuario string) bool{
	if _, existe := cuarto.usuarios[nombreUsuario]; existe{
		delete(cuarto.usuarios, nombreUsuario)
	}

	if len(cuarto.usuarios) == 0{
		return true
	}

	return false
}
