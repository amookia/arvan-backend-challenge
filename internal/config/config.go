package config

type (
	Config struct {
		Redis Redis `yaml:"redis"`
	}

	Redis struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
)
