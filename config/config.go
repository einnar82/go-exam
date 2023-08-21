package config

var (
	AuthToken          string
	StorageFolderPath  string
	DBConnectionString string
)

func LoadConfig() {
	AuthToken = "qGwnrOqsk0v1GYaeUTMH66yxvpFPSdCS"
	StorageFolderPath = "./disk/"
	DBConnectionString = "host=localhost port=5432 user=myuser password=mypassword dbname=mydb sslmode=disable"
}
