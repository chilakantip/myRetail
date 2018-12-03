package pg_persist

import (
	"fmt"
)

func init() {
	cfg := Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "",
		Password: "",
		Database: "my_retail_db",
	}
	if err := ConnectToPGDB(cfg); err != nil {
		fmt.Println("fail to get DB connection")
		fmt.Println(err)
	}
}
