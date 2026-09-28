using System.Text.Json;
using System.Text;
using cliente;
using mensaje;
using vista;

namespace controlador{

    ///<summary>
    ///Controlador del Cliente el cual se encargará de coordinar las
    ///acciones entre el Cliente (Modelo) y la vista en terminal (Vista).
    ///</summary>
    public class Controlador{
	private bool identificado;
	private Cliente cliente;
	private VistaCliente vista;

	///<summary>
	///Constructor del Controlador.
	///</summary>
	///<param name="cliente">Cliente que manejará.
	///<param name="vista">Vista que manejará.
	public Controlador(Cliente cliente, VistaCliente vista){
	    identificado = false;
	    this.cliente = cliente;
	    this.vista = vista;
	}

	///<summary>
	///Transforma un Mensaje a un JSON.
	///</summary>
	///<param name="msg">Mensaje a transformar.
	///<returns>Cadena con el JSON.
	public string MensajeAJSON(Mensaje msg){
	    if(msg == null) throw new ArgumentNullException(nameof(msg));

	    return JsonSerializer.Serialize(msg);
	}

	///<summary>
	///Transforma un JSON un Mensaje.
	///</summary>
	///<param name="datosJson">Cadena con el JSON.
	///<returns>Mensaje obtenido del JSON.
	public Mensaje? MensajeSinJSON(string datosJson){
	    return JsonSerializer.Deserialize<Mensaje>(datosJson);
	}

	///<summary>
	///Lee lo que escriba el usuario en la vista (terminal).
	///</summary>
	///<returns>Cadena de texto con lo que escribió el usuario.
	public string? LeerEntradaUsuario(){
	    return vista.LeerMensaje();
	}

	///<summary>
	///Ciclo en el que se escuchará continuamente al servidor para
	///recibir los mensajes que mande.
	///</summary>
	public void EscucharServidor(){
	    if(!cliente.Conectado())
		return;

	    string? msgServidor = cliente.Leer();

	    if(msgServidor != null){
		Mensaje? msg = MensajeSinJSON(msgServidor);

		if(msg != null){
		    ProcesaMensajeServidor(msg);
		}

		EscucharServidor();
	    }
	}

	///<summary>
	///Ciclo en el que se estará leyendo de forma continua lo que
	///escriba el usuario. 
	///</summary>
	public void LeerUsuario(){
	    if(!cliente.Conectado())
		return;

	    string? msg = LeerEntradaUsuario();

	    if(msg != null){
		ProcesaMensajeVista(msg);
	    }else{
		vista.EscribirMensaje("SISTEMA", "No puede escribir mensajes nulos.");
	    }

	    if(cliente.Conectado())
		LeerUsuario();
	}

	///<summary>
	///Maneja el mensaje que mandó el usuario y realiza las operaciones
	///necesarias de acuerdo a lo que solicita. 
	///</summary>
	///<param name="msg">Cadena de texto con el texto enviado por el usuario.
	public void ProcesaMensajeVista(string msg){
	    if(string.IsNullOrWhiteSpace(msg))
		return;
	    
	    Lector comando = new Lector();
	    msg = msg.Trim();
	    string[] msgSeparado = msg.Split(' ', StringSplitOptions.RemoveEmptyEntries);
	    bool valido = comando.ProcesaComando(msgSeparado);

	    if(!valido){
		vista.EscribirMensaje("SISTEMA", "Comando no válido.");
		return;
	    }

	    switch(msgSeparado[0]){
		case "/status":
		    string nuevoStatus = msgSeparado[1].Trim();
		    Mensaje cambiarStatus = Mensaje.CrearMensajeStatus(nuevoStatus);
		    string cambiarStatusJSON = MensajeAJSON(cambiarStatus);

		    cliente.EnviarDatos(cambiarStatusJSON);

		    vista.AgregaInicio();

		    break;
		    
		case "/list":
		    Mensaje listarUsuarios = Mensaje.CrearMensajeUsers();
		    string listarUsuariosJSON = MensajeAJSON(listarUsuarios);

		    cliente.EnviarDatos(listarUsuariosJSON);

		    break;

		case "/say":
		    string mensaje = msg.Substring(5);
		    Mensaje publicText = Mensaje.CrearMensajePublicText(mensaje);
		    string publicTextJSON = MensajeAJSON(publicText);

		    cliente.EnviarDatos(publicTextJSON);

		    vista.AgregaInicio();

		    break;

		case "/tell":
		    string usuarioDestino = msgSeparado[1];
		    string mensajeTell = msg.Substring(7 + usuarioDestino.Length);

		    Mensaje text = Mensaje.CrearMensajeText(usuarioDestino, mensajeTell);
		    string textJSON = MensajeAJSON(text);

		    cliente.EnviarDatos(textJSON);

		    vista.AgregaInicio();

		    break;
		    

		case "/createR":
		    string nombreCuarto = msgSeparado[1];
		    Mensaje cuarto = Mensaje.CrearMensajeNewRoom(nombreCuarto);
		    string cuartoJSON = MensajeAJSON(cuarto);

		    cliente.EnviarDatos(cuartoJSON);

		    break;

		case "/addR":
		    string nombreCuartoInvitacion = msgSeparado[1];
		    string[] invitados = msgSeparado[2].Split(',');
		    List<string> invitadosLista = new List<string>(invitados);

		    Mensaje invite = Mensaje.CrearMensajeInvite(nombreCuartoInvitacion, invitadosLista);
		    string inviteJSON = MensajeAJSON(invite);

		    cliente.EnviarDatos(inviteJSON);

		    vista.AgregaInicio();

		    break;

		case "/joinR":
		    string nombreCuartoAceptado = msgSeparado[1];

		    Mensaje aceptar = Mensaje.CrearMensajeJoinRoom(nombreCuartoAceptado);
		    string aceptarJSON = MensajeAJSON(aceptar);

		    cliente.EnviarDatos(aceptarJSON);

		    break;

		case "/listR":
		    string nombreCuartoLista = msgSeparado[1];

		    Mensaje listaCuarto = Mensaje.CrearMensajeRoomUsers(nombreCuartoLista);
		    string listaCuartoJSON = MensajeAJSON(listaCuarto);

		    cliente.EnviarDatos(listaCuartoJSON);

		    break;

		case "/sayR":
		    string nombreCuartoSay = msgSeparado[1];
		    string mensajeSayR = msg.Substring(7 + nombreCuartoSay.Length);

		    Mensaje roomText = Mensaje.CrearMensajeRoomText(nombreCuartoSay, mensajeSayR);
		    string roomTextJSON = MensajeAJSON(roomText);

		    cliente.EnviarDatos(roomTextJSON);

		    vista.AgregaInicio();

		    break;

		case "/leaveR":
		    string nombreCuartoLeave = msgSeparado[1];

		    Mensaje leaveRoom = Mensaje.CrearMensajeLeaveRoom(nombreCuartoLeave);
		    string leaveRoomJSON = MensajeAJSON(leaveRoom);

		    cliente.EnviarDatos(leaveRoomJSON);

		    vista.AgregaInicio();

		    break;

		case "/quit":
		    DesconectarCliente();
		    break;
	    }
	}

	///<summary>
	///Maneja el mensaje recibido por parte del servidor para transmitirlo 
	///al usuario.
	///</summary>
	///<param name="msg">Mensaje enviado por el servidor.
	public void ProcesaMensajeServidor(Mensaje msg){
	    if(msg == null || !msg.esValido()){
		return;
	    }

	    switch (msg.tipo){
		case "RESPONSE":
		    ProcesaResponse(msg);
		    break;

		case "NEW_USER":
		    vista.EscribirMensaje("SISTEMA", msg.username + " se ha conectado.");
		    break;

		case "NEW_STATUS":
		    vista.EscribirMensaje("SISTEMA", msg.username + " cambió su status a " + msg.status);
		    break;

		case "USER_LIST":
		    if(msg.users != null){
			StringBuilder respuesta = new StringBuilder("Lista de usuarios: ");

			foreach (KeyValuePair<string, string> usuario in msg.users){
			    respuesta.Append($"\n\t{usuario.Key} - {usuario.Value}");
			}
		    
			vista.EscribirMensaje("SISTEMA", respuesta.ToString());
		    }else{
			vista.EscribirMensaje("SISTEMA", "No hay usuarios.");
		    }
		    break;

		case "TEXT_FROM":
		    vista.EscribirMensajePrivado(msg?.username, msg?.text);
		    break;

		case "PUBLIC_TEXT_FROM":
		    vista.EscribirMensaje(msg?.username, msg?.text);
		    break;

		case "INVITATION":
		    vista.EscribirMensajeCuarto("SISTEMA", msg.username + " te ha invitado al cuarto " + msg.roomname);
		    break;

		case "JOINED_ROOM":
		    vista.EscribirMensajeCuarto(msg.roomname, msg.username + " se ha unido");
		    break;

		case "ROOM_USER_LIST":
		    if(msg.users != null){
			StringBuilder respuesta = new StringBuilder("Lista de usuarios: ");

			foreach (KeyValuePair<string, string> usuario in msg.users){
			    respuesta.Append($"\n\t{usuario.Key} - {usuario.Value}");
			}
		    
			vista.EscribirMensajeCuarto(msg.roomname, respuesta.ToString());
		    }else{
			vista.EscribirMensajeCuarto(msg.roomname, "No hay usuarios.");
		    }
		    break;

		case "ROOM_TEXT_FROM":
		    vista.EscribirMensajeCuarto(msg.username, msg.text);
		    break;

		case "LEFT_ROOM":
		    vista.EscribirMensajeCuarto(msg.roomname, msg.username + " ha abandonado el cuarto");
		    break;

		case "DISCONNECTED":
		    vista.EscribirMensaje("SISTEMA", msg.username + " se ha desconectado.");
		    break;
		    
		default:
		    Console.WriteLine($"se recibió un tipo distinto a response: {msg.tipo}.");
		    break;
	    }
	}

	///<summary>
	///Método privado que ayuda a manejar los distintos casos de los
	///Response enviados por el servidor.
	///</summary>
	///<param name="msg">Mensaje enviado por el servidor.
	private void ProcesaResponse(Mensaje msg){
	    switch(msg.operation){
		case "IDENTIFY":
		    if(msg.result == "SUCCESS"){
			identificado = true;
			vista.EscribirMensaje("SISTEMA", "Identificación exitosa. ¡Bienvenido!");
		    }else{
			vista.EscribirMensaje("SISTEMA", "Error de identificación: " + msg.result);
		    }
		    break;

		case "USER_LIST":
		    vista.EscribirMensaje("SISTEMA", "Se proporciona la lista de usuarios.");
		    break;
		    
		case "TEXT":
		    vista.EscribirMensaje("SISTEMA", "No se encontró el usuario de destino: " + msg.extra);
		    break;

		case "NEW_ROOM":
		    if(msg.result == "SUCCESS"){
			vista.EscribirMensajeCuarto("SISTEMA", "Se creó exitosamente el cuarto " + msg.extra);
		    }else if(msg.result == "ROOM_ALREADY_EXISTS"){
			vista.EscribirMensajeCuarto("SISTEMA", "El nombre del cuarto " + msg.extra + " ya está en uso.");
		    }
		    break;

		case "INVITE":
		case "JOIN_ROOM":
		case "ROOM_USERS":
		case "LEAVE_ROOM":
		case "ROOM_TEXT":
		    if(msg.result == "NO_SUCH_ROOM"){
			vista.EscribirMensajeCuarto("SISTEMA",  "El cuarto " + msg.extra + " no existe.");
		    }else if(msg.result == "NO_SUCH_USER"){
			vista.EscribirMensajeCuarto("SISTEMA", "No existe el usuario " + msg.extra);
		    }else if(msg.result == "NOT_INVITED"){
			vista.EscribirMensajeCuarto("SISTEMA", "No has sido previamente invitado al cuarto " + msg.extra);
		    }else if(msg.result == "NOT_JOINED"){
			vista.EscribirMensajeCuarto("SISTEMA", "No has sido invitado o no has aceptado la invitación al cuarto " + msg.extra);
		    }else if(msg.result == "SUCCESS"){
			vista.EscribirMensajeCuarto("SISTEMA", "Bienvenid@ a " + msg.extra + "!");
		    }
		    break;
		    
		case "INVALID":
		    if(msg.result == "NOT_IDENTIFIED"){
			vista.EscribirMensaje("SISTEMA", "Debe identificarse primero. Se le va a desconectar del sistema.");

			DesconectarCliente();
		    }

		    vista.EscribirMensaje("SISTEMA", "El mensaje está incompleto, con valores innesperados o no se puede reconocer.");
		    break;
	    }
	}

	///<summary>
	///Identifica al usuario con el servidor.
	///</summary>
	///<param name="username">Nombre de usuario con el que se quiere
	///identificar el usuario
	///<returns>True si el usuario se identificó exitosamente. False
	///si no se pudo identificar.
	public bool IdentificarCliente(string username){
	    Mensaje msgIdentify = Mensaje.CrearMensajeIdentify(username);
	    string identifyJSON = MensajeAJSON(msgIdentify);

	    cliente.EnviarDatos(identifyJSON);

	    string? respuestaJSON = cliente.Leer();

	    if(respuestaJSON == null){
		vista.EscribirMensaje("SERVIDOR", "Sin respuesta.");
		return false;
	    }

	    Mensaje? respuesta = MensajeSinJSON(respuestaJSON);

	    if(respuesta != null){
		ProcesaMensajeServidor(respuesta);
	    }

	    return identificado;
	}

	///<summary>
	///Desconecta al usuario del servidor.
	///</summary>
	public void DesconectarCliente(){
	    Mensaje msg = Mensaje.CrearMensajeDisconnect();
	    string msgJSON = MensajeAJSON(msg);

	    cliente.EnviarDatos(msgJSON);
	    cliente.Desconectar();
	}
    }
}
