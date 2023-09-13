package mongodb

import (
	"bases31/storage"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	db *mongo.Client
}

func InitDb() *Storage {

	db, err := connectToMongoDB()
	if err != nil {
		fmt.Println("Ошибка при подключении к MongoDB:", err)
	}
	storage := Storage{
		db: db,
	}
	return &storage
}

func (s *Storage) Posts() ([]storage.Post, error) {
	database := s.db.Database("news")
	collection := database.Collection("news")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	// Преобразуем результат в срез байтов с форматом JSON.
	var results []storage.Post
	if err := cursor.All(context.Background(), &results); err != nil {
		log.Fatal(err)
	}
	return results, nil
}

func (s *Storage) AddPost(stor storage.Post) error {
	collection := s.db.Database("news").Collection("news")
	_, err := collection.InsertOne(context.TODO(), stor)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("BD Mongo DB : Данные добавлены")
	return nil
}

func (s *Storage) UpdatePost(stor storage.Post) error {
	collection := s.db.Database("news").Collection("news")
	filter := bson.M{"id": stor.ID}
	update := bson.M{
		"$set": bson.M{"title": "Правим статью", "content": "Запросы SQL в pgAdmin"},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("BD Mongo DB : Данные обновлены")
	return nil
}

func (s *Storage) DeletePost(stor storage.Post) error {
	collection := s.db.Database("news").Collection("news")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"id": stor.ID})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("BD Mongo DB : Данные удалены")
	return nil
}

func connectToMongoDB() (*mongo.Client, error) {
	// Указываем параметры подключения к MongoDB.
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Подключаемся к MongoDB.
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Проверяем подключение к MongoDB.
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("Успешное подключение к MongoDB")
	return client, nil
}
