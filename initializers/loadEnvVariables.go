package initializers

import (
	"github.com/gookit/slog"
	"github.com/joho/godotenv"
)

// LoadEnvVariables
//
//	@Description: 加载.env环境变量
func LoadEnvVariables() {
	err := godotenv.Load("../.env")
	if err != nil {
		slog.Fatal("Error loading .env file")
		return
	}
	slog.Info("Env variables loaded")
}
