package mg_persist

import (
	"fmt"
	"testing"
)

func init() {
	cfg := Config{
		Host:       "localhost",
		Port:       "27017",
		Database:   "admin",
		Collection: "my_retail",
		User:       "",
		Password:   "",
	}
	if err := ConnectToMongoDB(cfg); err != nil {
		fmt.Println("fail to get DB connection")
		fmt.Println(err)
	}
}

//func TestGetNextQuestion(t *testing.T) {
//	fmt.Printf(Db.Name())
//	res, err := Db.Collection("my_retail").InsertOne(context.Background(),
//		bson.M{"ProductID": 10, "Name": "Kindle", "Description": "Amazon kindle Reader kit", "Type": "gadgets", "CreatedOn": time.Now()})
//	if err != nil {
//		t.Fatal(err)
//	}
//	id := res.InsertedID
//	fmt.Println(id)
//}

func TestFind(t *testing.T) {
	err := DeleteProduct(1600)
	if err != nil {
		t.Fatal(err)
	}
}

//func TestFind(t *testing.T) {
//	err := AddProduct(12, "test", "test", "test")
//	t.Fatal(err)
//}
