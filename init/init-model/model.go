package init_model

import (
	"github.com/jmoiron/sqlx"
	botgolang "github.com/mail-ru-im/bot-golang"
	"simplevkteamnotifiction/configs"
)

type Setting struct {
	Tables configs.Tables
	DB     *sqlx.DB
	VKTeam *botgolang.Bot
}
