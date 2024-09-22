package settings

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvMongo() string {
	//ex: "mongodb://foo:bar@localhost:27017"
	godotenv.Load(".env")
	return fmt.Sprintf("mongodb://%v:%v@%v:%v", getenv("DBUSER", "root"),
		getenv("DBPASS", "example"),
		getenv("DBHOST", "localhost"),
		getenv("DBPORT", "27017"))
}

func getenv(key string, alrt string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return alrt
}
