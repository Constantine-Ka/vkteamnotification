package model

type UserInfo struct {
	UserId          int    `json:"user_id"`
	UserLogin       string `json:"user_login"`
	UserName        string `json:"user_name"`
	UserEmail       string `json:"user_email"`
	UserGroup       string `json:"user_group"`
	UserBlocked     int    `json:"user_blocked"`
	MessagerAccount string `json:"messager_account"`
	IsBotAdd        string `json:"is_bot_add"`
}
