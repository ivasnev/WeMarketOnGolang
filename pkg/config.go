package pkg

import (
	"github.com/spf13/viper"
	"log"
)

// Config структура конфигурации
type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	DB struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
	} `mapstructure:"db"`
}

// LoadConfig загружает конфигурацию из файла и окружения
func LoadConfig(path string) (*Config, error) {
	// Установить автоматическое считывание переменных окружения
	viper.SetEnvPrefix("WEMARKET") // Префикс для переменных окружения
	viper.AutomaticEnv()           // Автоматически брать переменные из окружения

	// Привязка ключей к переменным окружения
	_ = viper.BindEnv("server.port", "WEMARKET_SERVER_PORT")
	_ = viper.BindEnv("db.user", "WEMARKET_DB_USER")
	_ = viper.BindEnv("db.password", "WEMARKET_DB_PASSWORD")
	_ = viper.BindEnv("db.name", "WEMARKET_DB_NAME")
	_ = viper.BindEnv("db.host", "WEMARKET_DB_HOST")
	_ = viper.BindEnv("db.port", "WEMARKET_DB_PORT")

	// Чтение файла конфигурации
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Ошибка чтения конфигурации из файла: %v", err)
	}

	// Заполнение структуры конфигурации
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func GetDBUrl(config *Config) string {
	return "host=" + config.DB.Host +
		" user=" + config.DB.User +
		" password=" + config.DB.Password +
		" dbname=" + config.DB.Name +
		" port=" + config.DB.Port +
		" sslmode=disable"

}
