package redis

type Config struct {
	Host     string
	Password string // no password set
	DB       int    // use default DB
}
