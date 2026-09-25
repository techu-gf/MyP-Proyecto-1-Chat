using System;
using System.Text;

namespace vista{

    ///<summary>
    ///VistaCliente manejará la lectura de lo que el usuario escribe
    ///en la terminal y lo que se escribe en la pantalla junto con cómo
    ///se escribe. 
    ///</summary>
    public class VistaCliente{

	private string textoTerminal = "";
	private readonly object lockEscritura = new object();

	///<summary>
	///Constructor de VistaCliente.
	///</summary>
	public VistaCliente(){ }

	///<summary>
	///Escribe el mensaje indicando quien lo envía y qué mensaje
	///quiere transmitir.
	///</summary>
	///<param name="sender">Quien manda el mensaje.
	///<param name="msg">Mensaje que se quiere mandar.
	public void EscribirMensaje(string? sender, string? msg){
	    lock(lockEscritura){
		int posicionActual = Console.CursorTop;
		Console.SetCursorPosition(0, posicionActual);
		Console.Write(new string(' ', Console.WindowWidth - 1));
		Console.SetCursorPosition(0, posicionActual);

		Console.WriteLine($"<{sender}> : {msg}");
		
		Console.Write("> " + textoTerminal);
	    }
	}

	public void EscribirMensajePrivado(string? sender, string? msg){
	    lock(lockEscritura){
		int posicionActual = Console.CursorTop;
		Console.SetCursorPosition(0, posicionActual);
		Console.Write(new string(' ', Console.WindowWidth - 1));
		Console.SetCursorPosition(0, posicionActual);

		Console.ForegroundColor = ConsoleColor.Yellow;
		Console.WriteLine($"<{sender}> : {msg}");
		Console.ResetColor();
		
		Console.Write("> " + textoTerminal);
	    }
	}

	///<summary>
	///Lee lo que escriba el usuario en la terminal.
	///</summary>
	///<returns>Cadena con el texto escrito en la terminal.
	public string? LeerMensaje(){
	    textoTerminal = "";

	    while(true){
		if(Console.KeyAvailable){
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
		
		}else{
		    Thread.Sleep(10);
		}
	    }
	}

	///<summary>
	///Agrega la flecha indicando el inicio del mensaje.
	///Más que nada es para que la terminal se vea más
	///homogenea. 
	///</summary>
	public void AgregaInicio(){
	    Console.Write("> " + textoTerminal);
	}
    }
}
