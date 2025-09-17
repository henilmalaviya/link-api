package entities

import (
	"context"
	"time"

	"api.link.henil.dev/internal/db"
	"api.link.henil.dev/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func generateWorkspaceSecret() string {
	secret, err := utils.GenerateRandomString(32)
	if err != nil {
		panic(err)
	}
	return secret
}

type Workspace struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	Name   string `bson:"name" json:"name"`
	Secret string `bson:"secret" json:"secret"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

func (w *Workspace) Insert(ctx context.Context) (primitive.ObjectID, error) {
	w.CreatedAt = time.Now().UTC()
	return Insert(ctx, db.GetWorkspaceCollection(), w)
}

func NewWorkspaceByName(name string) *Workspace {
	return &Workspace{
		Name:   name,
		Secret: generateWorkspaceSecret(),
	}
}

func GetWorkspaceByID(ctx context.Context, id string) (*Workspace, error) {
	return GetByID[Workspace](ctx, db.GetWorkspaceCollection(), id)
}

func GetWorkspaceByIDAndSecret(ctx context.Context, id, secret string) (*Workspace, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := map[string]any{
		"_id":    objID,
		"secret": secret,
	}

	return GetOneWithFilter[Workspace](ctx, db.GetWorkspaceCollection(), filter)
}

func GetWorkspaceBySecret(ctx context.Context, secret string) (*Workspace, error) {
	filter := map[string]any{
		"secret": secret,
	}
	return GetOneWithFilter[Workspace](ctx, db.GetWorkspaceCollection(), filter)
}
