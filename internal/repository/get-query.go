package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"simplevkteamnotifiction/configs"
	init_model "simplevkteamnotifiction/init/init-model"
	"simplevkteamnotifiction/internal/model"
)

func GetUsers(cfg init_model.Setting, logger *zap.Logger, id int) ([]model.UserInfo, error) {
	var UserList []model.UserInfo
	query := fmt.Sprintf("SELECT user_id, user_login, user_name, user_email, user_group, user_blocked, messager_account, is_bot_add FROM %s WHERE user_blocked=0 AND user_id=%d", cfg.Tables.Userlist, id)
	err := cfg.DB.Select(&UserList, query)
	if err != nil {
		logger.Error("Ошибка при получении списка пользователей по ID")
		return nil, err
	}
	return UserList, nil
}
func GetUsersChatID(db *sqlx.DB, cfg configs.Tables, logger *zap.Logger, id int) ([]string, error) {
	var UserList []string
	query := fmt.Sprintf("SELECT messager_account FROM %s WHERE user_blocked=0 AND user_id=%d", cfg.Userlist, id)
	err := db.Select(&UserList, query)
	if err != nil {
		logger.Info(query)
		logger.Error("Ошибка при получении списка пользователей по ID")
		return nil, err
	}
	return UserList, nil
}
