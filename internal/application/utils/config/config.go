package config

import "time"

//Config config
type Config struct {
	MysqlDSN      string        `mapstructure:"MYSQL_DSN"`
	RedisHost     string        `mapstructure:"REDIS_HOST"`
	AccessSecret  string        `mapstructure:"ACCESS_SECRET"`
	RefreshSecret string        `mapstructure:"REFRESH_SECRET"`
	Timeout       time.Duration `mapstructure:"TIMEOUT"`
	EndpointURL   string        `mapstructure:"EndpointUrl"`
	Region        string        `mapstructure:"Region"`
	ID            string        `mapstructure:"ID"`
	SecretKey     string        `mapstructure:"SecretKey"`
	AccessKey     string        `mapstructure:"AccessKey"`
	Profile       string        `mapstructure:"Profile"`
}
