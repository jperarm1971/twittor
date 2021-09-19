package main

import (
	"log"

	"github.com/jperarm1971/twittor/bd"
	"github.com/jperarm1971/twittor/handlers"
)

func main() {
	if bd.ChequeoConexion() == 0 {
		log.Fatal("Sin conexion a BBDD")
		return
	}

	handlers.Manejadores()

}
