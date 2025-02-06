package auth

import (
	"fmt"
	"log"
	"regexp"
	"testing"
)

func TestEye(t *testing.T) {
	type cat struct {
		Name  string
		Money int `json:"money" eye:"finance"`
	}
	vv := cat{Money: 1000}
	filter(&vv, "finance")

	t.Log(vv.Money)

}

func TestRegex(t *testing.T) {
	reg := regexp.MustCompile(fmt.Sprintf(`%s/([^/]+)([^?]*)`, "/api/shop"))
	if reg == nil { //解释失败，返回nil
		log.Print("1")
		return
	}
	match := reg.FindStringSubmatch("/api/shop/depot-order/202206241636319592613250/submit")
	if len(match) == 0 {
		log.Print("2")
		return
	}

	log.Println("match", match)
}
