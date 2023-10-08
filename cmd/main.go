package main

import (
	"database/sql"
	"errors"
	"fmt"
	botgolang "github.com/mail-ru-im/bot-golang"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path/filepath"
	"simplevkteamnotifiction/configs"
	"simplevkteamnotifiction/init/dbConnect"
	"simplevkteamnotifiction/init/init-model"
	"simplevkteamnotifiction/init/logger"
	"simplevkteamnotifiction/internal/repository"
	"simplevkteamnotifiction/internal/terminal"
)

func main() {
	terminal.GetFlags(SendText)
}

func SendText(c *cli.Context) error {
	//Инициализация модулей
	var setting init_model.Setting

	logg := logger.InitLogger()
	ex, err := os.Executable() //Определяем путь к запускаемому файлу
	if err != nil {
		logg.Error(fmt.Sprintf("%v", err))
		return errors.New(fmt.Sprintf("%v", err))
	}
	exPath := filepath.Dir(ex) // Определяем папку
	config := configs.Init(exPath, logg)
	DB, err := dbConnect.Init(config.DB)
	if err != nil {
		logg.Error(fmt.Sprintf("%v", err))
		return errors.New(fmt.Sprintf("%v", err))
	}
	setting.VKTeam, err = botgolang.NewBot(config.Vkbot.Token)
	if err != nil {
		logg.Error(fmt.Sprintf("%v", err))
		return errors.New(fmt.Sprintf("%v", err))
	}
	//Получение входящих данных из флагов
	recipientID := c.IntSlice("recipientID")
	recipientMail := c.StringSlice("recipient")
	messageInput := c.String("message")
	projectName := c.String("project")
	typeEvent := c.String("type")
	fileRoute := c.String("file")
	isFile := fileRoute != "."
	if len(recipientMail) == 0 && len(recipientID) == 0 {
		logg.Fatal("Получатель сообщения неуказан")
		return errors.New("Получатель сообщения неуказан")
	}
	if len(recipientID) > 0 {
		for _, id := range recipientID {
			users, err := repository.GetUsersChatID(DB, config.Tables, logg, id)
			if err != nil && errors.Is(err, sql.ErrNoRows) {
				logg.Warn(fmt.Sprintf("Пользователь с ID=%d не найден или заблокирован", id))
			} else if err != nil {
				logg.Error(fmt.Sprintf("Ошибка при поиске по id=%d| %v", id, err))
			}
			recipientMail = append(recipientMail, users...)
		}
	}
	for _, staff := range recipientMail {
		message := setting.VKTeam.NewMessage(staff)
		if isFile {
			file, err := os.OpenFile(fileRoute, os.O_RDWR|os.O_CREATE, 0755)
			if err != nil {
				logg.Error(fmt.Sprintf("Ошибка при открытии файла %s: %v", fileRoute, err))
			}
			message.AttachNewFile(file)
		}
		message.ParseMode = botgolang.ParseModeMarkdownV2
		message.Text = fmt.Sprintf("*%s|%s* \n%s", typeEvent, projectName, messageInput)
		errorSend := message.Send()
		if errorSend != nil {
			log.Print(errorSend)
			return errorSend
		}
	}
	defer logg.Sync() // Отложенная синхронизация записи логов

	return nil
}

func main0() {
	//var setting init_model.Setting
	//setting.Logger = *logger.InitLogger()
	//defer setting.Logger.Sync() // Отложенная синхронизация записи логов
	//
	//ex, err := os.Executable() //Определяем путь к запускаемому файлу
	//if err != nil {
	//	setting.Logger.Error(fmt.Sprintf("%v", err))
	//}
	//exPath := filepath.Dir(ex) // Определяем папку
	//config := configs.Init(exPath, &setting.Logger)
	//setting.DB, err = dbConnect.Init(config.DB)
	//setting.VKTeam, err = botgolang.NewBot(config.Vkbot.Token)
	//api.Server(db, config.Tables, logger)

}
