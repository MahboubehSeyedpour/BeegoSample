package main

import (
	_ "beegoSample/controllers"
	_ "beegoSample/routers"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq" // PostgreSQL driver
	"os"
)

func init() {

	// Load database configurations from app.conf
	dbUser := os.Getenv("DB_USER") // Since this field is stored as environment variable, it should be retrieved as such.
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName, _ := web.AppConfig.String("database::dbname")
	dbHost, _ := web.AppConfig.String("database::host")
	dbPort, _ := web.AppConfig.String("database::port")
	dbSSLMode, _ := web.AppConfig.String("database::sslmode")

	// Register PostgreSQL driver
	orm.RegisterDriver("postgres", orm.DRPostgres)

	// Build the database connection string
	dbConn := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName +
		" host=" + dbHost + " port=" + dbPort + " sslmode=" + dbSSLMode

	// Register database connection
	err := orm.RegisterDataBase("default", "postgres", dbConn)
	if err != nil {
		logs.Error("Database registration error:", err)
	}

	// Enable SQL debugging
	orm.Debug = true

	// Sync database
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Error("Database sync error:", err)
	}

	// Enable logging
	logs.SetLogger(logs.AdapterConsole)
	logs.SetLevel(logs.LevelDebug) // Log all debug/info/warning/errors
}

func main() {
	web.Run()
}
