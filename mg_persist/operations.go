package mg_persist

import (
	"context"
	"fmt"
	"time"

	"github.com/chilakantip/my_retail/env"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/pkg/errors"
)

var (
	ErrNoRecords      = errors.New("no_records")
	ErrNoRowsAffected = errors.New("no_rows_affected")
)

type ProductDetails struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	CreatedOn   string `json:"created_on"`
}

func GetProductDetails(id int) (prod *ProductDetails, err error) {
	result := bson.Raw{}
	filter := bson.D{{"ProductID", id}}
	err = Db.Collection(env.DBmgCollection).FindOne(context.Background(), filter).Decode(&result)
	if err != nil && len(result) > 0 {
		return nil, errors.Wrap(err, "failed to get product details")
	}
	if len(result) == 0 {
		return nil, ErrNoRecords
	}

	prod = &ProductDetails{}
	prod.Name = result.Lookup("Name").String()
	prod.Description, prod.Type = result.Lookup("Description").String(), result.Lookup("Type").String()
	prod.CreatedOn = result.Lookup("CreatedOn").String()

	return
}

func AddProduct(id int64, name, desc, type_ string) (err error) {
	res, err := Db.Collection(env.DBmgCollection).InsertOne(context.Background(), bson.M{
		"ProductID":   id,
		"Name":        name,
		"Description": desc,
		"Type":        type_,
		"CreatedOn":   time.Now().String()})
	if err != nil || res.InsertedID == nil {
		return errors.Wrap(err, "failed to add the product")
	}

	return
}

func UpdateProduct(id int, name, desc, type_ string) (err error) {
	result := Db.Collection(env.DBmgCollection).FindOneAndUpdate(context.Background(),
		bson.M{"ProductID": id},
		bson.M{"$set": bson.M{"Name": name, "Description": desc, "Type": type_}})

	if result.Err() != nil {
		return errors.Wrap(err, "failed to update the product")

	}

	return nil
}

func DeleteProduct(id int) (err error) {
	result, err := Db.Collection(env.DBmgCollection).DeleteOne(context.Background(), bson.M{"ProductID": id})
	fmt.Println(err)
	fmt.Println(result)
	if err != nil {
		return errors.Wrap(err, "failed to delete the product")

	}
	if result.DeletedCount == 0 {
		return ErrNoRowsAffected
	}

	return nil
}
