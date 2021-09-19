package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jperarm1971/twittor/models"
)

func GeneroJWT(t models.Usuario) (string, error) {
	miclave := []byte("MastersdelDesarrollo_grupodeFacebook")
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimientos,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miclave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
