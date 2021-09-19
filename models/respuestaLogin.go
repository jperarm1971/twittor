package models

type RespuestaLogin struct {
	Token string `bson:"token" json:"token,omitempty"`
}
