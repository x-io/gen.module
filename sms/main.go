package sms

import (
	"log"

	"github.com/x-io/gen.module/sms/adapter/wangjian"
	"github.com/x-io/gen.module/sms/adapter/yunxin"
	"github.com/x-io/gen.module/sms/core"
)

var _default string
var _debug bool

var adapter map[string]core.Adapter

type Config = core.Config

func Init(c Config) error {
	_debug = c.Debug
	_default = c.Default

	adapter = make(map[string]core.Adapter)

	for k, v := range c.Plugin {
		switch k {
		case "wangjian":
			adapter[k] = wangjian.New(v)
		case "yunxin":
			adapter[k] = yunxin.New(v)
		}
	}

	return nil
}

func SetDebug(debug bool) {
	_debug = debug
}

// Send sendSms
func Send(phone string, text string) (int, error) {
	if _debug {
		log.Println("debug", phone, text)
		return -1, nil
	}

	return adapter[_default].Send(phone, text)
}

// Send sendSms
func SendByChannel(channel, phone, text string) (int, error) {
	if _debug {
		log.Println("debug", phone, text)
		return -1, nil
	}

	if s, ok := adapter[channel]; ok {
		return s.Send(phone, text)
	}

	return adapter[_default].Send(phone, text)
}
