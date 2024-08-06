package toml

type Config struct {
	App struct {
		Info struct {
			Name    string
			Version string
		}
		Debug    bool
		Redis    map[string]RedisConf
		Mysql    map[string]MysqlConf
		RabbitMQ map[string]RabbitMQConf
	}
}
type RedisConf struct {
	Host     string
	Port     int
	Password string
	Database int
}
type MysqlConf struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}
type RabbitMQConf struct {
	Host     string
	Port     int
	User     string
	Password string
	Vhost    string
	Topic    map[string]string
	Queue    map[string]string
}
