using System;
using System.Text;

namespace vista{

    public class VistaCliente{

	private string textoTerminal = "";
	private readonly object lockEscritura = new object();

	public VistaCliente(){ }

	public void EscribirMensaje(string sender, string msg){
	    lock(lockEscritura){
		int posicionActual = Console.CursorTop;
		Console.SetCursorPosition(0, posicionActual);
		Console.Write(new string(' ', Console.WindowWidth - 1));
		Console.SetCursorPosition(0, posicionActual);

		Console.WriteLine($"<{sender}> : {msg}");
		
		Console.Write("> " + textoTerminal);
	    }
	}

	public string? LeerMensaje(){
	    textoTerminal = "";

	    while(true){
		ConsoleKeyInfo tecla = Console.ReadKey(true);

		lock(lockEscritura){
		    if(tecla.Key == ConsoleKey.Enter){
			Console.WriteLine();

			string mensaje = textoTerminal;
			textoTerminal = "";

			return mensaje;
		    }else if(tecla.Key == ConsoleKey.Backspace){
			if(textoTerminal.Length > 0){
			    textoTerminal = textoTerminal.Substring(0, textoTerminal.Length - 1);

			    Console.Write("\b \b");
			}
		    }else if(!char.IsControl(tecla.KeyChar)){
			textoTerminal += tecla.KeyChar;

			Console.Write(tecla.KeyChar);
		    }
		}
	    }
	}
    }
}
