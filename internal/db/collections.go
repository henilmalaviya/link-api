package db

import "go.mongodb.org/mongo-driver/mongo"

const workspaceCollectionName string = "workspaces"
const apiKeyCollectionName string = "api_keys"
const linkCollectionName string = "links"
const engagementCollectionName string = "engagements"

func GetWorkspaceCollection() *mongo.Collection {
	return getDB().Collection(workspaceCollectionName)
}

func GetApiKeyCollection() *mongo.Collection {
	return getDB().Collection(apiKeyCollectionName)
}

func GetLinkCollection() *mongo.Collection {
	return getDB().Collection(linkCollectionName)
}

func GetEngagementCollection() *mongo.Collection {
	return getDB().Collection(engagementCollectionName)
}
