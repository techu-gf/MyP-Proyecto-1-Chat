using cliente;
using System;
using vista;
using controlador;
using System.Threading;

class EjecutaCliente{

    static async Task Main(string[] args){
	Lector banderas = new Lector();
	banderas.ProcesaArgs(args);

	Console.WriteLine($"Conectando a {banderas.GetHost()}:{banderas.GetPuerto()}.\n");

	try{
	    Cliente cliente = new Cliente(banderas.GetPuerto(), banderas.GetHost(), banderas.GetUsername());
	    VistaCliente vista = new VistaCliente();
	    Controlador ctrl = new Controlador(cliente, vista);

	    if(cliente.Conectado()){
		Console.WriteLine("Usuario conectado, seguimos con la revisión del username.\n");

		bool identificado = ctrl.IdentificarCliente(banderas.GetUsername());

		if(!identificado){
		    cliente.Desconectar();
		    return;
		}

		Thread hiloServidor = new Thread(ctrl.EscucharServidor);
		hiloServidor.IsBackground = true;
		hiloServidor.Start();

		ctrl.LeerUsuario();
	    }
	}catch (System.Net.Sockets.SocketException e){
	    Console.WriteLine($"Error al conectar con el servidor: {e.Message}");
	}catch (Exception e){
	    Console.WriteLine($"Ocurrió un error inesperado: {e.Message}");
	}
    }
}
