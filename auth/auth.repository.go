package auth

import (
	"context"

	"github.com/mnshah219/go_net_http/auth/dto"
	"github.com/mnshah219/go_net_http/auth/schema"
	"github.com/mnshah219/go_net_http/db"
	"go.mongodb.org/mongo-driver/mongo"
)

func getCollection() *mongo.Collection {
	client := db.GetClient()
	return client.Collection(schema.USER_TABLE)
}
func createUser(user dto.SignupDto) (*mongo.InsertOneResult, error) {
	coll := getCollection()
	result, err := coll.InsertOne(context.TODO(), user)
	return result, err
}

func findOneUser(filter interface{}) schema.User {
	coll := getCollection()
	result := coll.FindOne(context.TODO(), filter)
	user := schema.User{}
	result.Decode(&user)
	return user
}
