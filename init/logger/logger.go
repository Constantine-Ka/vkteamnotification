package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

func InitLogger() *zap.Logger {
	os.MkdirAll("/var/log/vkteam", 0666)
	os.OpenFile("/var/log/vkteam/app.log", os.O_WRONLY, 0666)
	os.OpenFile("/var/log/vkteam/error.log", os.O_WRONLY, 0666)
	cfg := zap.Config{
		Encoding:         "json",                                // Формат логов: json или console
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),   // Уровень логирования: Debug, Info, Warn, Error, DPanic, Panic или Fatal
		OutputPaths:      []string{"/var/log/vkteam/app.log"},   // Путь к файлу для записи логов
		ErrorOutputPaths: []string{"/var/log/vkteam/error.log"}, // Путь к файлу для записи ошибок
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",                   // Ключ для сообщения
			LevelKey:    "level",                     // Ключ для уровня логирования
			TimeKey:     "time",                      // Ключ для времени
			EncodeTime:  zapcore.ISO8601TimeEncoder,  // Формат времени: ISO8601 или EpochMillis
			EncodeLevel: zapcore.CapitalLevelEncoder, // Формат уровня логирования: CapitalCase или Lowercase
		},
	}
	logger, err := cfg.Build()
	if err != nil {
		log.Fatal("Failed to create logger:", err)
	}

	return logger
}
