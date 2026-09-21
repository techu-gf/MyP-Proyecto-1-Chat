using cliente;
using System;

class EjecutaCliente{

    static async Task Main(string[] args){
	Bandera banderas = new Bandera(args);

	Console.WriteLine($"Conectando A {banderas.GetHost()}:{banderas.GetPuerto()}");

	try{
	    Cliente cliente = new Cliente(banderas.GetPuerto(), banderas.GetHost(), banderas.GetUsername());

	    if(cliente.Conectado()){
		Console.WriteLine("Conexión exitosa con el servidor.");

		string? respuesta = Console.ReadLine();

		if(respuesta == "SALIR")
		    cliente.Desconectar();
	    }
	}catch (System.Net.Sockets.SocketException e){
            Console.WriteLine($"Error al conectar con el servidor: {e.Message}");
            Console.WriteLine("Asegúrate de que el servidor Go esté ejecutándose y el puerto sea correcto.");
	}catch (Exception e){
            Console.WriteLine($"Ocurrió un error inesperado: {e.Message}");
	}
    }
}
