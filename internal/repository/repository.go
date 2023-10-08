package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

func Auth(db *sqlx.DB, tablename, login, password string) int {
	var data int
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE user_blocked=0 AND user_group LIKE '%%admin%%'AND user_login='%s' AND user_plain_pass='%s';", tablename, login, password)
	err := db.Get(&data, query)
	if err != nil && err != sql.ErrNoRows {
		log.Println(query)
		log.Print(err)
		return data
	} else if err == sql.ErrNoRows {
		return 0
	}
	return data
}
