package config

type (
	Config struct {
		Redis      Redis      `yaml:"redis"`
		Middleware Middleware `yaml:"middleware"`
		Mongo      Mongo      `yaml:"mongo"`
	}

	Redis struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}

	Mongo struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		Name string `yaml:"name"`
	}

	Middleware struct {
		UserKey   string `yaml:"userkey"`
		PerMinute int    `yaml:"perminute"`
		Monthly   int    `yaml:"monthly"`
	}
)
