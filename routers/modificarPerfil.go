package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jperarm1971/twittor/bd"
	"github.com/jperarm1971/twittor/models"
)

/* funcion para grabar en bd el usuario*/

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	var status bool
	status, err = bd.ModificaRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "No se pudo modificar el registro"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha podido modificar el perfil del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
