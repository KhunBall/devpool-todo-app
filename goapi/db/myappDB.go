package db

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// const dbName_sqlite = "C:\\Users\\ball_\\Desktop\\Devpool_Final\\goapi\\db\\myapp.db"
// const dbName_mssql = "sqlserver://sa:abcABC123@localhost:1433?database=master"

// "mssql", "sqlserver://username:password@localhost:1433?

var isMigrate = false

func DB() *gorm.DB {

	// os.Setenv("USER", "sa")
	// os.Setenv("PASSWORD", "abcABC123")
	// os.Setenv("HOST", "db")
	// os.Setenv("PORT", "1433")
	// os.Setenv("DATABASE", "master")
	
	dbName_mssql := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("DATABASE"),
	)
	// db, err := gorm.Open(sqlite.Open(dbName_sqlite), &gorm.Config{})
	db, err := gorm.Open(sqlserver.Open(dbName_mssql), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	return db
}

func Migrate() {
	db := DB()
	db.AutoMigrate(&UserDB{})
}
