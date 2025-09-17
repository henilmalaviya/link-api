package entities

import (
	"context"
	"fmt"
	"time"

	"api.link.henil.dev/internal/db"
	"api.link.henil.dev/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ApiKeyPermission string

const (
	ApiKeyPermissionLinkRead       ApiKeyPermission = "link:read"
	ApiKeyPermissionLinkWrite      ApiKeyPermission = "link:write"
	ApiKeyPermissionEngagementRead ApiKeyPermission = "engagement:read"
)

func generateApiKey() string {
	key, err := utils.GenerateRandomString(16)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("key_%s", key)
}

type ApiKey struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	Key   string `bson:"key" json:"key"`
	Label string `bson:"label" json:"label"`

	Permissions []ApiKeyPermission `bson:"permissions" json:"permissions"`

	WorkspaceID primitive.ObjectID `bson:"workspace_id" json:"workspace_id"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

func (w *ApiKey) Insert(ctx context.Context) (primitive.ObjectID, error) {
	w.CreatedAt = time.Now().UTC()
	return Insert(ctx, db.GetApiKeyCollection(), w)
}

func (w *ApiKey) GetByID(ctx context.Context) (*ApiKey, error) {
	return GetByID[ApiKey](ctx, db.GetApiKeyCollection(), w.ID.Hex())
}

func NewApiKey(label string, permissions []ApiKeyPermission, workspaceId primitive.ObjectID) *ApiKey {
	return &ApiKey{
		Key:         generateApiKey(),
		Label:       label,
		Permissions: permissions,
		WorkspaceID: workspaceId,
	}
}

func NewApiKeyById(id string) (*ApiKey, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return &ApiKey{ID: objId}, nil
}
