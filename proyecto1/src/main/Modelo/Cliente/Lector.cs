using System;

namespace cliente{

    ///<summary>
    ///Lector nos ayudará a leer banderas o identificar el comando
    ///dado en una cadena de texto que dividiremos en arreglos. 
    ///</summary>
    public class Lector{

	private int puerto = 8080;
	private string username = string.Empty;
	private string host = "127.0.0.1";

	///<summary>
	///Constructor del Lector.
	///</summary>
	public Lector(){}

	///<summary>
	///Regresa el puerto predeterminado o el obtenido por las banderas.
	///</summary>
	///<returns>Puerto indicado en las banderas.
	public int GetPuerto(){
	    return puerto;
	}

	///<summary>
	///Regresa el nombre de usuario obtenido por las banderas.
	///</summary>
	///</returns>Nombre de usuario indicado en las banderas.
	public string GetUsername(){
	    return username;
	}

	///<summary>
	///Regresa la IP obtenido por las banderas.
	///</summary>
	///</returns>IP indicado en las banderas.
	public string GetHost(){
	    return host;
	}

	///<summary>
	///Procesa el arreglo de cadenas buscando el puerto, el host y el username.
	///</summary>
	///<param name="args">Arreglo de cadenas a procesar.
	public void ProcesaArgs(string[] args){
	    bool hayPuerto = false;
	    bool hayUsername = false;
	    bool hayHost = false;

	    for(int i = 0; i < args.Length; i++){
		switch(args[i]){
		    case "-p":
			if(hayPuerto){
			    Console.WriteLine("Solo puede ingresar un puerto.");
			    Environment.Exit(1);
			}

			try{
			    if(i + 1 >= args.Length){
				Console.WriteLine("Debe dar un puerto.");
				Environment.Exit(1);
			    }
			
			    puerto = int.Parse(args[i+1]);
			    i++;
			    hayPuerto = true;
			}catch (FormatException ex){
			    Console.WriteLine($"Ingrese un puerto válido. Error: {ex}");
			    Environment.Exit(1);
			}
			break;
		    
		    case "-u":
			if(hayUsername){
			    Console.WriteLine("Solo puede ingresar un nombre de usuario.");
			    Environment.Exit(1);
			}
		    
			if(i + 1 >= args.Length){
			    Console.WriteLine("Debe dar un usuario.");
			    Environment.Exit(1);
			}
		    
			username = args[i+1];
			i++;
			hayUsername = true;
			break;
		    
		    case "-h":
			if(hayHost){
			    Console.WriteLine("Solo puede ingresar un host.");
			    Environment.Exit(1);
			}

			if(i + 1 >= args.Length){
			    Console.WriteLine("Debe dar un host.");
			    Environment.Exit(1);
			}
		    
			host = args[i+1];
			i++;
			hayHost = true;
			break;

		    default:
			break;
		}
	    }
	    

	    if(!hayPuerto){
		Console.WriteLine("Debe proporcionar un puerto\n");
		Environment.Exit(1);
	    }

	    if(!hayUsername){
		Console.WriteLine("Debe proporcional un nombre de usuario.\n");
		Environment.Exit(1);
	    } 
	}

	///<summary>
	///Procesa el arreglo de cadenas buscando identificar el tipo de
	///comando que tiene y si cumple con las características necesarias.
	///</summary>
	///<param name="comando">Arreglo de cadenas a procesar.
	///<returns>True si es un comando válido y cumple con lo necesario.
	///False en otros casos.
	public bool ProcesaComando(string[] comando){
	    if(comando == null || comando.Length == 0){
		return false;
	    }
	    
	    switch(comando[0]){
		case "/list":
		case "/quit":
		    return true;

		case "/status":
		    if(comando.Length < 2){
			return false;
		    }
		    return true;
		    
		case "/say":
		    if(comando.Length < 2){
			return false;
		    }
		    return true;
		    
		default:
		    return false;
	    }
	}
    }
}
