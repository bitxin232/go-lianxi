package config

type Config struct {
	DatabaseDSN   string
	ServerAddress string
}

func Load() *Config {
	return &Config{
		DatabaseDSN:   "root:123@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local",
		ServerAddress: ":8080",
	}
}
