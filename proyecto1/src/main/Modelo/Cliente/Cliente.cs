using System.IO;
using System.Net.Sockets;
using System.Threading;
using mensaje;

namespace cliente{

    ///<summary>
    ///Clase que realizará la conexión con el servidor. Contiene el nombre
    ///de usuario seleccionado, un indicador de si está conectado o no al servidor 
    ///y demás atributos necesarios para la conexión con el servidor.
    ///</summary>
    public class Cliente{
	private int puerto;
	private string host;
	private string username;
	private TcpClient? socket;
	private StreamWriter? writer;
	private StreamReader? reader;
	private bool conectado;

	///<summary>
	///Constructor del cliente. Realiza la conexión con el servidor.
	///</summary>
	///<param name="puerto">Número de puerto donde está el servido.
	///<param name="host">IP donde está el servidor.
	///<param name="username">Como desea identificarse el usuario en 
	///el servidor.
	public Cliente(int puerto,string host, string username){
	    this.puerto = puerto;
	    this.host = host;
	    this.username = username;
	    conectado = false;

	    try{
	        socket = new TcpClient(host, puerto);
		NetworkStream stream = socket.GetStream();
	    
		reader = new StreamReader(stream);
		writer = new StreamWriter(stream) {AutoFlush = true};
		conectado = true;
	    }catch(SocketException e){
		Console.WriteLine($"Servidor no encendido o puerto cerrado: {e}");
		throw;
	    }
	}

	///<summary>
	///Indica si el cliente está conectado al servidor.
	///</summary>
	///<returns>True si está conectado. False si no lo está.
	public bool Conectado(){
	    return conectado;
	}

	///<summary>
	///Recibe los mensajes del servidor.
	///</summary>
	///<returns>Cadena con el mensaje del servidor.
	public string? Leer(){
	    try{
		return reader?.ReadLine();
	    }catch{
		return null;
	    }
	}

	///<summary>
	///Envía mensajes JSON al servidor.
	///</summary>
	///<param name="json">JSON que se enviará al servidor.
	public void EnviarDatos(string json){
	    if(string.IsNullOrEmpty(json)){
		Console.WriteLine("No se pueden mandar datos vacíos.\n");
		return;
	    }

	    if(!conectado){
		Console.WriteLine("No está identificado con el servidor.\n");
		Environment.Exit(1);
		
	    }
		writer?.WriteLine(json);
	}

	///<summary>
	///Nos desconecta del servidor.
	///</summary>
	public void Desconectar(){
	    if(!conectado)
		return;
	    
	    conectado = false;
	    reader?.Close();
	    writer?.Close();
	    socket?.Close();
	}
    }
}
