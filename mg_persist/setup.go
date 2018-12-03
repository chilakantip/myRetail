package mg_persist

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var Db *mongo.Database

type Config struct {
	Host       string
	Port       string
	User       string
	Password   string
	Database   string
	Collection string
}

func ConnectToMongoDB(cfg Config) (err error) {
	//"mongodb://127.0.0.1:27017"
	client, err := mongo.NewClient(fmt.Sprintf("mongodb://%s:%s@%s:%s", cfg.User, cfg.Password, cfg.Host, cfg.Port))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	Db = client.Database(cfg.Database)

	return
}
