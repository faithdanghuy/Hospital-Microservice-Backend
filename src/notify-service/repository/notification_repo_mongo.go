package repository

import (
	"context"
	"time"

	"github.com/Hospital-Microservice/notify-service/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoNotificationRepo struct {
	col *mongo.Collection
}

func NewMongoNotificationRepo(col *mongo.Collection) *mongoNotificationRepo {
	return &mongoNotificationRepo{col: col}
}

func (r *mongoNotificationRepo) Save(ctx context.Context, n *entity.NotificationEntity) error {
	if n.ID.IsZero() {
		n.ID = primitive.NewObjectID()
	}
	n.IsRead = false
	n.CreatedAt = time.Now().UTC()
	_, err := r.col.InsertOne(ctx, n)

	return err
}

func (r *mongoNotificationRepo) ListByUser(ctx context.Context, userID string, limit, offset int64) ([]entity.NotificationEntity, error) {
	filter := bson.M{"user_id": userID}
	opts := options.Find().SetSkip(offset).SetLimit(limit).SetSort(bson.D{{Key: "created_at", Value: -1}})
	cur, err := r.col.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	var out []entity.NotificationEntity
	if err := cur.All(ctx, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (r *mongoNotificationRepo) MarkRead(ctx context.Context, id string) error {
	_, err := r.col.UpdateByID(ctx, id, bson.M{
		"$set": bson.M{
			"is_read": true,
			"read_at": time.Now().UTC(),
		},
	})
	return err
}
