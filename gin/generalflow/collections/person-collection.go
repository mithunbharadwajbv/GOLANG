package collections

import (
	"context"
	"time"

	"github.com/mithun/generalflow/config"
	"github.com/mithun/generalflow/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

type PersonCollection interface {
	SavePerson(entity.Person)
}

type personCollection struct {
	collection *mongo.Collection
}

func NewPersonCollection() PersonCollection {
	var db *mongo.Database = config.MongoConnect()
	return &personCollection{
		collection: db.Collection("Person"),
	}
}

func (collection *personCollection) SavePerson(person entity.Person) {
	collection.collection.InsertOne(ctx, person)
}
