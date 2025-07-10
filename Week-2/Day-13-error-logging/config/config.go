package config

type Config struct {
	Port  int
	DBDSN string
	Env   string
}

var AppConfig Config

func Load() {
	AppConfig = Config{}
}
