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

	public void ProcesaMensajeServidor(Mensaje msg){
	    if(msg == null || !msg.esValido()){
		return;
	    }

	    switch (msg.tipo){
		case "RESPONSE":
		    ProcesaResponse(msg);
		    break;
		default:
		    Console.WriteLine("se recibió un tipo distinto a response.");
		    break;
	    }
	}

	private void ProcesaResponse(Mensaje msg){
	    switch(msg.operation){
		case "IDENTIFY":
		    if(msg.result == "SUCCESS"){
			identificado = true;
			vista.EscribirMensaje("SISTEMA", "Identificación exitosa.");
		    }else{
			vista.EscribirMensaje($"SISTEMA", "Error de identificación ({msg.result}).");
		    }
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

	public void OperacionInvalida(){
	    
	}
    }
}
