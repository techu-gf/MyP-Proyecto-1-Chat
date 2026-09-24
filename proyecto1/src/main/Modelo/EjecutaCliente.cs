using cliente;
using System;
using vista;
using controlador;
using System.Threading;

class EjecutaCliente{
     
    static void Main(string[] args){
	Lector banderas = new Lector();
	banderas.ProcesaArgs(args);

	try{
	    Cliente cliente = new Cliente(banderas.GetPuerto(), banderas.GetHost(), banderas.GetUsername());
	    VistaCliente vista = new VistaCliente();
	    Controlador ctrl = new Controlador(cliente, vista);

	    if(cliente.Conectado()){
		bool identificado = ctrl.IdentificarCliente(banderas.GetUsername());

		if(!identificado){
		    cliente.Desconectar();
		    return;
		}
		
		Thread hiloServidor = new Thread(ctrl.EscucharServidor);
		hiloServidor.IsBackground = true;
		hiloServidor.Start();

		Thread.Sleep(100);

		ctrl.LeerUsuario();
	    }
	}catch (System.Net.Sockets.SocketException e){
	    Console.WriteLine($"Error al conectar con el servidor: {e.Message}");
	}catch (Exception e){
	    Console.WriteLine($"Ocurrió un error inesperado: {e.Message}");
	}
    }
}
