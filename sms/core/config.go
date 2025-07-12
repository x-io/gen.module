package core

type Config struct {
	Debug   bool
	Default string
	Plugin  map[string]Engine
}

type Engine struct {
	Domain string
	Uid    string
	Key    string
	Sign   string
}

// Adapter Cache适配器接口
type Adapter interface {
	Send(phone, content string) (int, error)
}
