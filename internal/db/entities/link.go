package entities

import (
	"context"
	"time"

	"api.link.henil.dev/internal/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Link struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	Title *string `bson:"title" json:"title"`
	Url   string  `bson:"url" json:"url"`

	ShortName *string `bson:"short_name" json:"short_name"`
	Enabled   bool    `bson:"enabled" json:"enabled"`

	WorkspaceID primitive.ObjectID `bson:"workspace_id" json:"workspace_id"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func (l *Link) Insert(ctx context.Context) (primitive.ObjectID, error) {
	l.CreatedAt = time.Now().UTC()
	l.UpdatedAt = time.Now().UTC()
	return Insert(ctx, db.GetLinkCollection(), l)
}

func (l *Link) GetByIDOrShortName(ctx context.Context) (*Link, error) {
	return GetOneWithFilter[Link](ctx, db.GetLinkCollection(), map[string]any{
		"$or": []any{
			map[string]any{"_id": l.ID},
			map[string]any{"short_name": l.ShortName},
		},
	})
}

func NewLink(title *string, url string, shortName *string, workspaceID primitive.ObjectID) *Link {
	return &Link{
		Title:       title,
		Url:         url,
		ShortName:   shortName,
		Enabled:     true,
		WorkspaceID: workspaceID,
	}
}

func NewLinkByIdAndShortName(id string, shortName string) (*Link, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return &Link{
		ID:        objID,
		ShortName: &shortName,
	}, nil
}
