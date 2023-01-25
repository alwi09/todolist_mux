package configuration

type Config struct {
	DbUser     string `json:"dbUser" default:"alwi09"`
	DbPassword string `json:"dbPassword" default:"alwiirfani11"`
	DbHost     string `json:"dbHost" default:"localhost"`
	DbPort     int    `json:"dbPort" default:"1234"`
	DbName     string `json:"dbName" default:"rest_full_api_todolist"`
}
