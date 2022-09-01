package main

import (
	"fmt"
	"mongodb/database"
	"mongodb/service"
)

func main() {
	database.Connect()

	//bsonM := bson.M{"foo": "barM", "hello": "worldM", "pi": 3.14159}
	//bsonD := bson.D{{"foo", "barD"}, {"hello", "worldD"}, {"pi", 3.14159}}
	//bsonA := bson.A{"bar", "worldA", 3.14159, bson.D{{"qux", 12345}}}
	//res, err := collection.InsertOne(context.Background(), bsonM)
	//res, err = collection.InsertOne(ctx, bsonD)

	//service.CreateUser()
	users := service.GetAllUser()
	for i, s := range users {
		fmt.Println(i, s)
	}
	user := users[0]
	user1, _ := service.FindUserByEmail(user.Email)
	fmt.Println(user1)
	user2, _ := service.UpdateUserById(user.UserId)
	fmt.Println(user2)

	fmt.Println("Paginate Page")
	users, total := service.GetUsers(1, 5)
	for i, s := range users {
		fmt.Println(i, s)
	}
	fmt.Println(total)

	fmt.Println("sort user")
	user3 := service.SortingUserByAge()
	for i, s := range user3 {
		fmt.Println(i, s)
	}

}
