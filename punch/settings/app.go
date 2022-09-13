package settings

type AppConfig struct {
	Server *server
	Mysql  *mysql
	Redis  *redis
	Jwt    *jwt
}

type server struct {
	Port int `yaml:"port"`
}

type mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

type jwt struct {
	Key string `yaml:"key"`
}
