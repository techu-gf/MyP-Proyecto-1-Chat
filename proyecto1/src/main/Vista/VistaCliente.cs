using System;

namespace vista{

    public class VistaCliente{

	public VistaCliente(){}

	public void EscribirMensaje(string sender, string msg){
	    Console.WriteLine($"<{sender}> : {msg}\n");
	}

	public string? LeerMensaje(){
	    return Console.ReadLine();
	}
    }
}
