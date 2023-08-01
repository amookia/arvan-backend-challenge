package config

type (
	Config struct {
		Service    Service    `yaml:"service"`
		Redis      Redis      `yaml:"redis"`
		Middleware Middleware `yaml:"middleware"`
		Mongo      Mongo      `yaml:"mongo"`
	}

	Service struct {
		Port string `yaml:"port" envconfig:"SERVICE_PORT"`
	}

	Redis struct {
		Host string `yaml:"host" envconfig:"REDIS_HOST"`
		Port string `yaml:"port" envconfig:"REDIS_PORT"`
	}

	Mongo struct {
		Host string `yaml:"host" envconfig:"MONGO_HOST"`
		Port string `yaml:"port" envconfig:"MONGO_PORT"`
		Name string `yaml:"name" envconfig:"MONGO_DBNAME"`
	}

	Middleware struct {
		PerMinute int    `yaml:"perminute" default:"10"`
		Monthly   int    `yaml:"monthly"   default:"9000000" ` // in bytes
	}
)
