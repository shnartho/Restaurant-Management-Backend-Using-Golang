package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	ID               primitive.ObjectID `bson:"_id"`
	Number_of_guests *int               `bson:"number_of_guests" validate:"required"`
	Table_number     *int               `bson:"table_number" validate:"required"`
	Created_at       time.Time          `bson:"created_at"`
	Updated_at       time.Time          `bson:"updated_at"`
	Table_id         string             `bson:"table_id"`
}
