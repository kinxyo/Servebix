package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/kinxyo/Servebix.git/config"
	"log"
)

// cfg --> configuration

func GetDB() (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 config.Env.DBUser,
		Passwd:               config.Env.DBPassword,
		Addr:                 config.Env.DBAddress,
		DBName:               config.Env.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func DbInit(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[db] Database initialized...")
}
