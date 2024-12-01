package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConventToObjectId(id string) primitive.ObjectID {
	objectId, _ := primitive.ObjectIDFromHex(id)
	return objectId
}
