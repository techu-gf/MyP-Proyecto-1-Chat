# Proyecto de Chat - Modelado y Programación

## Lenguajes de Programación y Sistema de Construcción

Servidor - Golang
Cliente - C#

Para el sistema de construcción se usa Taskfile. Para su instalación se recomienda seguir la [documentación oficial](https://taskfile.dev/docs/installation).

Dentro de este proyecto tenemos 3 tasks:

* servidor: se encargará de la ejecución del servidor.
* cliente: se encargará de la ejecución del cliente.
* clean: hará la limpieza de archivos basura, temporales o binarios de compilación que se generen a partir de la ejecución del código. 

## Ejecución del Servidor

Posicionate en el directorio raíz del proyecto, justo donde se encuentra el archivo Taskfile y ejecuta la siguiente línea:

``` bash
task servidor -- -p <PUERTO>
```

donde <PUERTO> se sustituye por el puerto en el que se encontrará el servidor. En caso de simplemente escribir

```
task servidor
```
el valor predeterminado del puerto será 1234.

## Ejecución del Cliente

Posicionate en el directorio raíz del proyecto, justo donde se encuentra el archivo Taskfile y ejecuta la siguiente línea:

```
task cliente -- -p <PUERTO> -h <HOST> -u <USUARIO>
```

donde <PUERTO> se sustituye por el puerto en el que está el servidor, <HOST> se sustituye por la IP donde está el servidor y <USUARIO> será el nombre de usuario deseado. En caso de no especificar el host, el valor predeterminado es 127.0.0.1.

## Comandos

| Comando | Acción |
| :--- | :--- |
| /status <NUEVO_STATUS> | Cambia el status a BUSY, AWAY o ACTIVE. |
| /list | Enlista los usuarios en el servidor junto con sus status. |
| /say <MENSAJE> | Manda un mensaje a todos los usuarios. |
| /tell <USUARIO_DESTINO> <MENSAJE> | Manda un mensaje privado al usuario elegido. |
| /createR <NOMBRE_CUARTO> | Crea un cuarto con el nombre dado. |
| /addR <NOMBRE_INVITADO> | Invita al usuario deseado. |
| /joinR <NOMBRE_CUARTO> | Acepta la invitación, y se une, al cuarto elegido. |
| /listR <NOMBRE_CUARTO> | Enlista los usuarios en el cuarto escrito junto con sus status. |
| /sayR <NOMBRE_CUARTO> <MENSAJE> | Manda un mensaje a todos los usuarios del cuarto deseado. |
| /leaveR <NOMBRE_CUARTO> | Abandona el cuarto elegido. |
| /quit | Abandona el servidor. |