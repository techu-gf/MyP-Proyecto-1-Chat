using System.Text.Json;
using cliente;
using mensaje;
using vista;

namespace controlador{

    public class Controlador{
	private bool identificado;
	private Cliente cliente;
	private VistaCliente vista;
	
	public Controlador(Cliente cliente, VistaCliente vista){
	    identificado = false;
	    this.cliente = cliente;
	    this.vista = vista;
	}

	public string MensajeAJSON(Mensaje msg){
	    if(msg == null) throw new ArgumentNullException(nameof(msg));

	    return JsonSerializer.Serialize(msg);
	}

	public Mensaje? MensajeSinJSON(string datosJson){
	    return JsonSerializer.Deserialize<Mensaje>(datosJson);
	}

	public string? LeerEntradaUsuario(){
	    return vista.LeerMensaje();
	}

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

	public void ProcesaMensajeVista(string msg){
	    Lector comando = new Lector();
	    string[] msgSeparado = msg.Split(' ');
	    bool valido = comando.ProcesaComando(msgSeparado);

	    if(!valido){
		vista.EscribirMensaje("SISTEMA", "Comando no válido.");
		return;
	    }

	    switch(msgSeparado[0]){
		case "/list":
		    Mensaje listarUsuarios = Mensaje.CrearMensajeUsers();
		    string listarUsuariosJSON = MensajeAJSON(listarUsuarios);

		    cliente.EnviarDatos(listarUsuariosJSON);

		    break;

		case "/quit":
		    DesconectarCliente();
		    break;
	    }
	}

	public void ProcesaMensajeServidor(Mensaje msg){
	    if(msg == null || !msg.esValido()){
		return;
	    }

	    switch (msg.tipo){
		case "RESPONSE":
		    ProcesaResponse(msg);
		    break;
		    
		default:
		    Console.WriteLine($"se recibió un tipo distinto a response: {msg.tipo}.");
		    break;
	    }
	}

	private void ProcesaResponse(Mensaje msg){
	    switch(msg.operation){
		case "IDENTIFY":
		    if(msg.result == "SUCCESS"){
			identificado = true;
			vista.EscribirMensaje("SISTEMA", "Identificación exitosa. ¡Bienvenido!");
		    }else{
			vista.EscribirMensaje($"SISTEMA", "Error de identificación ({msg.result}).");
		    }
		    break;

		case "USER_LIST":
		    vista.EscribirMensaje("SISTEMA", "Se proporciona la lista de usuarios.");
		    break;
		    
	    }
	}

	public bool IdentificarCliente(string username){
	    Mensaje msgIdentify = Mensaje.CrearMensajeIdentify(username);
	    string identifyJSON = MensajeAJSON(msgIdentify);

	    cliente.EnviarDatos(identifyJSON);

	    string? respuestaJSON = cliente.Leer();

	    if(respuestaJSON == null){
		vista.EscribirMensaje("SERVIDOR", "Sin respuesta.\n");
		return false;
	    }

	    Mensaje? respuesta = MensajeSinJSON(respuestaJSON);

	    if(respuesta != null){
		ProcesaMensajeServidor(respuesta);
	    }

	    return identificado;
	}

	public void DesconectarCliente(){
	    Mensaje msg = Mensaje.CrearMensajeDisconnect();
	    string msgJSON = MensajeAJSON(msg);

	    cliente.EnviarDatos(msgJSON);
	    cliente.Desconectar();

	    vista.EscribirMensaje("SISTEMA", "Desconexión exitosa.");
	}

	public void OperacionInvalida(){
	    
	}
    }
}
