package vkteam

import (
	"errors"
	botgolang "github.com/mail-ru-im/bot-golang"
)

func SentToVkteamText(bot *botgolang.Bot, recipient, content string) (string, error) {
	var NotDenied = errors.New("error while sending text: error status from API: Permission denied")

	message := bot.NewTextMessage(recipient, content)
	err := message.Send()
	if err != nil {
		if err == NotDenied {
			return "", errors.New("Получатель не добавил бота")
		}
		return "", err
	}
	return message.ID, nil
}
