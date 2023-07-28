package object

import "github.com/google/uuid"

type (
	ObjectModel struct {
		CheckSum string    `bson:"checksum"`
		Size     int64     `bson:"size"`
		Owner    string    `bson:"owner"`
		Uuid     uuid.UUID `bson:"uuid"`
	}
)
