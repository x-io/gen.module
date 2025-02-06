package auth

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/x-io/gen"
	"github.com/x-io/gen.module/jwt"
	"github.com/x-io/gen/errors"
)

// Middleware
func Middleware(file string) gen.Middleware {
	//	option := prepareOptions(options)
	config := getConfig(file)
	return func(c *gen.Context) {
		meta := c.Meta("route")
		data := c.Data("JWT")

		if v, ok := data.(jwt.MapClaims); ok {
			acl := v["acl"].(string)
			eye := v["eye"].(string)

			if acl == "*" || acl == "+" {
				c.Next()
			} else if isAuth(config, c.Request.Method, c.Request.URL.Path, acl) || isSkip(meta) {
				c.Next()
				c.Result = filter(c.Result, eye)
			}
		}
		c.Write(errors.HTTP(http.StatusForbidden))

	}
}

func isAuth(config *Config, method, url, acl string) bool {
	reg := regexp.MustCompile(fmt.Sprintf(`%s/([^/]+)([^?]*)`, config.Prefix))
	if reg == nil { //解释失败，返回nil
		return false
	}
	match := reg.FindStringSubmatch(url)
	if len(match) == 0 {
		return false
	}

	if v, ok := config.Rule[match[1]]; ok {
		key := match[1]

		role, err := getRole(acl, key)
		if err != nil {
			return false
		}

		// log.Println(role, key, acl, match)
		switch method {
		case "GET":
			return true
		case "POST":
			return isEvent(role, "create", strings.Split(v, ","))
		case "DELETE":
			return isEvent(role, "delete", strings.Split(v, ","))
		case "PUT":
			return isEvent(role, getExt(match[2]), strings.Split(v, ","))
		}
	}

	return false
}

func isSkip(meta string) bool {
	return strings.Contains(meta+",", "^acl,") || strings.Contains(meta+",", "^jwt,")
}

func getRole(acl, key string) (int64, error) {
	reg := regexp.MustCompile(fmt.Sprintf(`,%s:([^,]+)`, key))
	match := reg.FindStringSubmatch("," + acl)
	if len(match) > 0 {
		if a, err := strconv.ParseInt(match[1], 16, 32); err == nil {
			return a, nil
		} else {
			log.Println(err)
		}
		return 0, nil
	}

	return 0, fmt.Errorf("error")
}

func isEvent(vv int64, key string, rules []string) bool {
	// log.Printf("vv:%d, key:%s, rule:%v", vv, key, rules)
	for i, v := range rules {
		if strings.Contains(v+"|", key+"|") {
			return vv>>i&0x01 == 1
		}
	}

	return false
}

func getExt(path string) string {
	if path == "" {
		return "change"
	}

	for i := len(path) - 1; i >= 1; i-- {
		if path[i] == '/' {
			return path[i+1:]
		}
	}

	if len(path) < 15 {
		return path[1:]
	}

	return "change"
}
