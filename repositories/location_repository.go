package repositories

import (
	"context"
	"log"

	"github.com/madjiebimaa/isyarah/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoLocationRepository struct {
	coll *mongo.Collection
}

func NewMongoLocationRepository() models.LocationRepository {
	return &mongoLocationRepository{}
}

func (m *mongoLocationRepository) Create(ctx context.Context, location *models.Location) error {
	_, err := m.coll.InsertOne(ctx, location)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
