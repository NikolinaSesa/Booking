package persistence

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient(host, port string) (*mongo.Client, error) {
	//uri := fmt.Sprintf("mongodb://%s:%s", host, port)
	options := options.Client().ApplyURI("mongodb+srv://draga:draga@cluster0.dlhjqkp.mongodb.net/?retryWrites=true&w=majority")
	return mongo.Connect(context.TODO(), options)

}
