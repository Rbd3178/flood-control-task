package limiter

// Config
type Config struct {
	RedisAddr string `toml:"redisaddr" json:"redisaddr"`
	RedisPass string `toml:"redispass" json:"redispass"`
	RedisDB   int    `toml:"redisdb" json:"redisdb"`
	Interval  int    `toml:"interval" json:"interval"`
	Limit     int64  `toml:"limit" json:"limit"`
}

// NewConfig
func NewConfig() *Config {
	return &Config{
		RedisAddr: "localhost:6379",
		RedisPass: "",
		RedisDB:   0,
		Interval:  1,
		Limit:     1,
	}
}
