package configs

type Yamlcfg struct {
	DB     DB     `yaml:"db"`
	Tables Tables `yaml:"tables"`
	Vkbot  Vkbot  `yaml:"vkbot"`
}

type DB struct {
	Driver   string `yaml:"driver"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Sslmode  string `yaml:"sslmode"`
}
type Tables struct {
	Userlist string `yaml:"userlist"`
}

type Vkbot struct {
	BotID int    `yaml:"botId"`
	Nick  string `yaml:"nick"`
	Token string `yaml:"token"`
}
