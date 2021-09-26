package routers

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jperarm1971/twittor/bd"
	"github.com/jperarm1971/twittor/models"
)

var Email string
var IDUsuario string

func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miclave := []byte("MastersdelDesarrollo_grupodeFacebook")

	claims := &models.Claim{}

	//splitToken := strings.Split(tk, "Bearer")
	//if len(splitToken) != 2 {
	//	return claims, false, string(""), errors.New("formato invalido de token")
	//}
	//tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miclave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token invalido")
	}
	return claims, false, string(""), err
}
