package middlew

import (
	"net/http"

	"github.com/jperarm1971/twittor/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConexion() == 0 {
			http.Error(w, "Se ha perdido  la conexion con la BBDD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
