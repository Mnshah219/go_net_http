package schema

var USER_TABLE string = "users"

type MongoDocument struct {
	ID string `bson:"_id"`
}
type User struct {
	Username      string
	Email         string
	Password      string
	MongoDocument `bson:"inline"`
}
