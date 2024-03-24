package limiter

// Config
type Config struct {
	RedisAddr string `toml:"redisaddr"`
	RedisPass string `toml:"redispass"`
	RedisDB   int    `toml:"redisdb"`
	Interval  int    `toml:"interval"`
	Limit     int64    `toml:"limit"`
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
