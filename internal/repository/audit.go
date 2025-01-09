package repository

import (
	"context"
	"os"

	audit "github.com/Dolg0ff/todoapp-audit-log/pkg/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type Audit struct {
	db *mongo.Database
}

func NewAudit(db *mongo.Database) *Audit {
	return &Audit{
		db: db,
	}
}

func (r *Audit) Insert(ctx context.Context, item audit.LogItem) error {
	_, err := r.db.Collection(os.Getenv("DB_COLLECTION")).InsertOne(ctx, item)

	return err
}
