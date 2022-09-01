package service

import (
	"context"
	"fmt"
	"log"
	"mongodb/database"
	"mongodb/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateUser() {
	collection := database.Collection("users")
	n := 0
	for n < 25 {
		user := model.NewUser()
		count, err := CheckUserEmailExist(user.Email)
		fmt.Println(count)
		if err == nil && count == 0 {
			_, err := collection.InsertOne(context.Background(), user)
			fmt.Println(err)
		}
		n++
	}
}

func UpdateUserById(id string) (model.User, error) {
	var user model.User
	userid, _ := primitive.ObjectIDFromHex(id)
	user = model.NewUser()
	user.ID = userid
	user.UserId = id
	collection := database.Collection("users")
	// Find the employee and update its data
	query := bson.D{{Key: "_id", Value: userid}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: user.Name},
				{Key: "email", Value: user.Email},
				{Key: "gender", Value: user.Gender},
				{Key: "phone", Value: user.Phone},
				{Key: "age", Value: user.Age},
			},
		},
	}
	err := collection.FindOneAndUpdate(context.Background(), query, update).Err()
	return user, err
}

func GetAllUser() []model.User {
	//var user model.User
	//ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	//defer cancel()
	collection := database.Collection("users")
	cursor, _ := collection.Find(context.Background(), bson.D{{}})

	//var users []primitive.M
	var users []model.User
	for cursor.Next(context.Background()) {
		//var user bson.M
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
		//fmt.Println(user)
	}

	return users
}

func SortingUserByAge() []model.User {
	//var user model.User
	//ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	//defer cancel()
	opts := options.Find()
	// Sort by `price` field descending
	opts.SetSort(bson.D{{"age", -1}})
	collection := database.Collection("users")
	cursor, _ := collection.Find(context.Background(), bson.D{{}}, opts)

	//var users []primitive.M
	var users []model.User
	for cursor.Next(context.Background()) {
		//var user bson.M
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
		//fmt.Println(user)
	}

	return users
}

func GetUsers(page, limit int) ([]model.User, int64) {
	collection := database.Collection("users")
	filter := bson.D{{}} // selects all documents
	options := new(options.FindOptions)
	if limit != 0 {
		if page == 0 {
			page = 1
		}
		options.SetSkip(int64((page - 1) * limit))
		options.SetLimit(int64(limit))
	}

	cursor, _ := collection.Find(context.Background(), filter, options)
	total, _ := collection.CountDocuments(context.Background(), filter)

	var users []model.User
	for cursor.Next(context.Background()) {
		//var user bson.M
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
		//fmt.Println(user)
	}

	return users, total
}

func FindUserByEmail(email string) (model.User, error) {
	var user model.User
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	collection := database.Collection("users")
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	return user, err
}

func CheckUserEmailExist(email string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	collection := database.Collection("users")
	count, err := collection.CountDocuments(ctx, bson.M{"email": email})
	return count, err
}
