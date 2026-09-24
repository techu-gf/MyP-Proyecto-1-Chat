using System;
using System.Text.Json;
using System.Text.Json.Serialization;
using System.Collections.Generic;

namespace mensaje{

    public class Mensaje{
	[JsonPropertyName("type")]
	[JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
	public string? tipo {get;set;}
	
	[JsonPropertyName("username")]
	[JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
	public string? username {get;set;}
	
	[JsonPropertyName("operation")]
	[JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
	public string? operation {get;set;}
	
	[JsonPropertyName("result")]
	[JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
	public string? result {get;set;}
	
	[JsonPropertyName("extra")]
	[JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
	public string? extra {get;set;}
	
	[JsonPropertyName("status")]
	[JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
	public string? status {get;set;}
	
	[JsonPropertyName("users")]
	[JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
	public Dictionary<string, string>? users {get;set;}
	
	[JsonPropertyName("usernames")]
	[JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
	public List<string>? usernames {get;set;}
	
	[JsonPropertyName("text")]
	[JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
	public string? text {get;set;}
	
	[JsonPropertyName("roomname")]
	[JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
	public string? roomname {get;set;}
    
	public static Mensaje CrearMensajeIdentify(string username){
	    return new Mensaje{tipo = "IDENTIFY", username = username};
	}

	public static Mensaje CrearMensajeResponse(string operation, string result, string extra){
	    return new Mensaje{tipo = "RESPONSE", operation = operation, result = result, extra = extra};
	}

	public static Mensaje CrearMensajeNewUser(string username){
	    return new Mensaje{tipo = "NEW_USER", username = username};
	}

	public static Mensaje CrearMensajeStatus(string status){
	    return new Mensaje{tipo = "STATUS", status = status};
	}

	public static Mensaje CrearMensajeNewStatus(string username, string status){
	    return new Mensaje{tipo = "NEW_STATUS", username = username, status = status};
	}

	public static Mensaje CrearMensajeUsers(){
	    return new Mensaje{tipo = "USERS"};
	}

	public static Mensaje CrearMensajeUserList(Dictionary<string, string> users){
	    return new Mensaje{tipo = "USER_LIST", users = users};
	}

	public static Mensaje CrearMensajeText(string username, string text){
	    return new Mensaje{tipo = "TEXT", username = username, text = text};
	}

	public static Mensaje CrearMensajeTextFrom(string username, string text){
	    return new Mensaje{tipo = "TEXT_FROM", username = username, text = text};
	}

	public static Mensaje CrearMensajePublicText(string text){
	    return new Mensaje {tipo = "PUBLIC_TEXT", text = text };
	}

	public static Mensaje CrearMensajePublicTextFrom(string username, string text){
	    return new Mensaje {tipo = "PUBLIC_TEXT_FROM", username = username, text = text };
	}

	public static Mensaje CrearMensajeNewRoom(string roomname){
	    return new Mensaje {tipo = "NEW_ROOM", roomname = roomname };
	}

	public static Mensaje CrearMensajeInvite(string roomname, List<string> usernames){
	    return new Mensaje {tipo = "INVITE", roomname = roomname, usernames = usernames };
	}

	public static Mensaje CrearMensajeInvitation(string username, string roomname){
	    return new Mensaje {tipo = "INVITATION", username = username, roomname = roomname };
	}

	public static Mensaje CrearMensajeJoinRoom(string roomname){
	    return new Mensaje {tipo = "JOIN_ROOM", roomname = roomname };
	}

	public static Mensaje CrearMensajeJoinedRoom(string roomname, string username){
	    return new Mensaje {tipo = "JOINED_ROOM", roomname = roomname, username = username };
	}

	public static Mensaje CrearMensajeRoomUsers(string roomname){
	    return new Mensaje {tipo = "ROOM_USERS", roomname = roomname };
	}

	public static Mensaje CrearMensajeRoomText(string roomname, string text){
	    return new Mensaje {tipo = "ROOM_TEXT", roomname = roomname, text = text };
	}

	public static Mensaje CrearMensajeRoomTextFrom(string roomname, string username, string text){
	    return new Mensaje {tipo = "ROOM_TEXT_FROM", roomname = roomname, username = username, text = text };
	}

	public static Mensaje CrearMensajeLeaveRoom(string roomname){
	    return new Mensaje {tipo = "LEAVE_ROOM", roomname = roomname };
	}

	public static Mensaje CrearMensajeLeftRoom(string roomname, string username){
	    return new Mensaje {tipo = "LEFT_ROOM", roomname = roomname, username = username };
	}

	public static Mensaje CrearMensajeDisconnect(){
	    return new Mensaje {tipo = "DISCONNECT" };
	}

	public static Mensaje CrearMensajeDisconnected(string username){
	    return new Mensaje {tipo = "DISCONNECTED", username = username };
	}

	public bool esValido(){
	    return !string.IsNullOrWhiteSpace(tipo);
	}
    }
}
    
