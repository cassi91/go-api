package config

const (
	DBDriver = "sqlite3"
	DBPath   = "./db.sqlite3"
)

func GetDbConnectionString() string {
	return DBPath
}
