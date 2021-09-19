package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jperarm1971/twittor/bd"
	"github.com/jperarm1971/twittor/models"
)

/* funcion para grabar en bd el usuario*/

func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "La contraseña debe tener al menos 6 caracteres", 400)
		return
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con este correo", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Se produjo en error en el guardado del registro en la BD:"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se pudo insertar el usuario en la BD:"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
