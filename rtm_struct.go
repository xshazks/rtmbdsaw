package rtmsaw

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Monitor struct {
	ID         			primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Proker				string             `bson:"proker" json:"proker"`
	Status 				string             `bson:"status" json:"status"`
	About 				string             `bson:"about" json:"about"`
	Karyawan 			string             `bson:"karyawan" json:"karyawan"`
}
