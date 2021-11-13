package apiserver

import "github.com/AnnDutova/database_course_work/internal/app/store"

type Config struct {
	BinAddr  string `toml:"bind_addr"` //адрес на кот запускаем сервер
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BinAddr:  ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
