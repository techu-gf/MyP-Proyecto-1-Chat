using System.IO;
using System.Net.Sockets;
using System.Threading;
using mensaje;

namespace cliente{

    public class Cliente{
	private int puerto;
	private string host;
	private string username;
	private string status;
	private TcpClient? socket;
	private StreamWriter? writer;
	private StreamReader? reader;
	private bool conectado;

	public Cliente(int puerto,string host, string username){
	    this.puerto = puerto;
	    this.host = host;
	    this.username = username;
	    status = "ACTIVE";
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

	public string GetStatus(){
	    return status;
	}
	
	public bool Conectado(){
	    return conectado;
	}

	public string? Leer(){
	    try{
		return reader?.ReadLine();
	    }catch{
		return null;
	    }
	}

	public void EnviarDatos(string json){
		writer?.WriteLine(json);
	}

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
