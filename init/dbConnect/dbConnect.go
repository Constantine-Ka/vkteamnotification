package dbConnect

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"simplevkteamnotifiction/configs"
)

func Init(cfg configs.DB) (*sqlx.DB, error) {
	var err error
	var db *sqlx.DB
	if cfg.Driver == "postgres" {
		db, err = sqlx.Open(cfg.Driver, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Dbname, cfg.Sslmode))
	} else {
		db, err = sqlx.Connect("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname))
	}
	if err != nil {
		log.Println()
		log.Fatal("Ошибка подключения к БД. ", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Ошибка при пинге БД. ", err)
		return nil, err
	}
	return db, nil
}
