package auth

import (
	"log"

	"github.com/x-io/gen.module/system"
)

type Config struct {
	Prefix string            `json:"prefix"`
	Rule   map[string]string `json:"rule"`
}

func getConfig(file string) *Config {
	var config Config
	if err := system.LoadConfig(&config, file, "/acl/"); err != nil {
		log.Printf("ERROR: load file:%s err:%s", file, err)
	}
	return &config
}
