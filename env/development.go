package env

func init() {

	AppName = "my_retail"
	Varsion = "v1.0.0"
	AppEnv = "development"
	ServiceOnPort = ":3000"

	DBHost = "localhost"
	DBPort = "5432"
	DBUser = "my_retail"
	DBPassword = "my_retail123"
	DBDatabase = "my_retail_db"

	DBmgHost = "localhost"
	DBmgPort = "27017"
	DBmgUser = ""
	DBmgPassword = ""
	DBmgDatabase = "admin"
	DBmgCollection = "my_retail"

	CurrencyType = "USD"

}
