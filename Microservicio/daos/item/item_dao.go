package item

import (
	"context"
	"fmt"
	"microservicio/utils/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
func GetById(id string) model.Book {
	var book model.Book
	db := db.MongoDb
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return book
	}
	err = db.Collection("books").FindOne(context.TODO(), bson.D{{"_id", objID}}).Decode(&book)
	if err != nil {
		fmt.Println(err)
		return book
	}
	return book

}

func Insert(book model.Book) model.Book {
	db := db.MongoDb
	insertBook := book
	insertBook.Id = primitive.NewObjectID()
	_, err := db.Collection("books").InsertOne(context.TODO(), &insertBook)

	if err != nil {
		fmt.Println(err)
		return book
	}
	book.Id = insertBook.Id
	return book
}
*/
func InsertItem(item model.item) model.item {
	db := db.MongoDb
	InsertItem := item
	InsertItem.Id = primitive.NewObjectID()
	_, err := db.Collection("items").ImsertOne(context.TODO(), &InsertItem)

	if err != nil {
		fmt.Println(err)
		return item
	}

	item.Id = InsertItem.Id
	return item
}
