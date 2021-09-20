package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Config *Config

	Client     *mongo.Client
	CurrentDB  *mongo.Database
	Collection *mongo.Collection
}

func NewDB() *DB {
	return &DB{
		Config: NewConfig(),
	}
}

//* Метод подключения к базе данных
func (db *DB) Open() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(db.Config.URI))
	if err != nil {
		return err
	}

	err = client.Connect(context.TODO())
	if err != nil {
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	db.Client = client
	db.CurrentDB = client.Database(db.Config.DatabaseName)
	db.Collection = db.CurrentDB.Collection(db.Config.CollectionName)

	return nil
}

//* Метод закрытия базы данных
func (db *DB) Close() error {
	err := db.Client.Disconnect(context.TODO())
	if err != nil {
		return err
	}

	return nil
}

//* Метод получения документа из базы данных по его ID
func (db *DB) Get(id string) (Respond, error) {
	var respond Respond

	err := db.Collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&respond)
	if err != nil {
		return Respond{}, err
	}

	return respond, nil
}

//* Метод создания документа в базе данных
func (db *DB) Create(id string) error {
	_, err := db.Collection.InsertOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}

//* Метод удаления документа из базы данных по его ID
func (db *DB) Delete(id string) error {
	_, err := db.Collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}

//* Метод обновления Refresh-токена документа по его ID
func (db *DB) Update(id, newVal string) error {
	_, err := db.Collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{
				{"Refresh-Token", newVal},
			},
			},
		},
	)
	if err != nil {
		return err
	}

	return nil
}
