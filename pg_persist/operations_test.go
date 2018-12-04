package pg_persist

import (
	"fmt"
	"testing"
)

func init() {
	cfg := Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "my_retail",
		Password: "my_retail123",
		Database: "my_retail_db",
	}
	if err := ConnectToPGDB(cfg); err != nil {
		fmt.Println("fail to get DB connection")
		fmt.Println(err)
	}
}

func TestAddProduct(t *testing.T) {
	err := DeleteProduct(2000)
	t.Fatal(err)

}
