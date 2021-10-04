package models

type GraboTweet struct {
	UserID  string `bson:"userid" json:"userid,omitempty"`
	Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   string `bson:"fecha" json:"fecha,omitempty"`
}
