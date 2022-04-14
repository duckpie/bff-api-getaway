package config

type Config struct {
	Services      ServicesConfigs      `toml:"services"`
	Microservices MicroservicesConfigs `toml:"microservices"`
}

type ServicesConfigs struct {
	Server ServerConfig `toml:"server"`
}

type ServerConfig struct {
	Host string `toml:"HOST"`
	Port int64  `toml:"PORT"`
}

type MicroservicesConfigs struct {
	UserMs     UserMsConfig     `toml:"user"`
	SecurityMs SecurityMsConfig `toml:"security"`
}

type UserMsConfig struct {
	Host string `toml:"HOST"`
	Port int64  `toml:"PORT"`
}

type SecurityMsConfig struct {
	Host string `toml:"HOST"`
	Port int64  `toml:"PORT"`
}

func NewConfig() *Config {
	return &Config{}
}
