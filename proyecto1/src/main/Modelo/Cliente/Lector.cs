using System;

namespace cliente{

    public class Lector{

	private int puerto = 8080;
	private string username = string.Empty;
	private string host = "127.0.0.1";
	private string tipoOperacion = string.Empty;

	public Lector(){}

	public int GetPuerto(){
	    return puerto;
	}

	public string GetUsername(){
	    return username;
	}

	public string GetHost(){
	    return host;
	}

	public string GetOperacion(){
	    return tipoOperacion;
	}

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

	public bool ProcesaComando(string[] comando){
	    if(comando == null || comando.Length == 0){
		return false;
	    }
	    
	    switch(comando[0]){
		case "/list":
		case "/quit":
		    return true;
		    
		case "/say ":
		    if(1 > comando.Length){
			Console.WriteLine("Debe proporcionar el mensaje.");
			return false;
		    }
		    return true;
		    
		default:
		    return false;
	    }
	}
    }
}
