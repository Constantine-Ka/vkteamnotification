package terminal

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func GetFlags(callback func(ctx *cli.Context) (string, error)) (string, error) {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntSliceFlag{
				Name:     "recipientID",
				Aliases:  []string{"id", "rid"},
				Required: false,
				Usage:    "ID (сотрудника) Получателя сообщения",
			},
			&cli.StringSliceFlag{
				Name:     "recipient",
				Aliases:  []string{"to", "r"},
				Required: false,
				Usage:    "Ник (или почта) Получателя сообщения",
			},
			&cli.StringFlag{
				Name:     "project",
				Value:    "Без проекта",
				Aliases:  []string{"p", "proj"},
				Required: false,
				Usage:    "Название проекта",
			},
			&cli.StringFlag{
				Name:    "type",
				Value:   "INFO",
				Aliases: []string{"t"},
				Usage:   "Тип сообщения (Информирование, предупреждение, Ошибка, Паника или просто задача)",
			},
			&cli.StringFlag{
				Name:     "message",
				Value:    "",
				Aliases:  []string{"m", "mess", "msg", "mes"},
				Required: true,
				Usage:    "Сообщение для пользователя",
			},
			&cli.StringFlag{
				Name:     "file",
				Value:    ".",
				Aliases:  []string{"f", "path", "filepath"},
				Usage:    "Необязателен. Указывается путь к файлу",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "check",
				Value:    false,
				Aliases:  []string{"chck", "chek", "chec"},
				Usage:    "Проверка возможности отправки сообщений. Если флаг указан, то сообщенние удалится сразу после отправки",
				Required: false,
			},
			&cli.BoolFlag{
				Name:    "keyboard",
				Aliases: []string{"k"},
				Usage:   "Бесполезен. Функционал пока не реализован",
			},
		},
		Action: func(context *cli.Context) error {
			msgID, err := callback(context)
			fmt.Println(msgID)
			os.Exit(0)
			return err
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	return "", nil
}
