package entities

import (
	"context"
	"time"

	"api.link.henil.dev/internal/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EngagementType string

const (
	EngagementTypeClick   EngagementType = "click"
	EngagementTypeQR      EngagementType = "qr"
	EngagementTypeUnknown EngagementType = "unknown"
)

type Engagement struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	LinkID primitive.ObjectID `bson:"link_id,omitempty" json:"link_id"`
	Type   EngagementType     `bson:"type" json:"type"`

	ShouldCount bool `bson:"should_count" json:"should_count"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

func (e *Engagement) Insert(ctx context.Context) (primitive.ObjectID, error) {
	return Insert(ctx, db.GetEngagementCollection(), e)
}

func NewEngagement(linkID primitive.ObjectID, engagementType EngagementType, shouldCount bool) *Engagement {
	return &Engagement{
		ID:          primitive.NewObjectID(),
		LinkID:      linkID,
		Type:        engagementType,
		ShouldCount: shouldCount,
		CreatedAt:   time.Now(),
	}
}
