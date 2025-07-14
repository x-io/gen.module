package yunxin

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/x-io/gen.module/sms/core"
)

type engine struct {
	core.Engine
}

func New(c core.Engine) *engine {
	return &engine{
		c,
	}
}

func (s *engine) Send(phone string, content string) (int, error) {

	data := make(url.Values)
	data["userCode"] = []string{s.Uid}
	data["userPass"] = []string{s.Key}
	data["DesNo"] = []string{phone}
	data["Msg"] = []string{s.Sign + content}

	// fmt.Println(s.Domain, data)
	res, err := http.PostForm(s.Domain, data)
	if err != nil {
		fmt.Println(err.Error())
		return -100, err
	}
	defer res.Body.Close()
	// fmt.Println("post send success")

	d, err := io.ReadAll(res.Body)
	if err != nil {
		return -100, err
	}

	result, ok := XMLToMap(string(d))["string"]
	if !ok {
		return -100, fmt.Errorf(string(d))
	}

	ret, err := strconv.Atoi(result.(string))
	if err != nil {
		return -1, err
	}

	msg, ok := YunxinErr[ret]
	if ok {
		return ret, fmt.Errorf(msg)
	}

	return 1, nil
}

var YunxinErr = map[int]string{
	-1:  "应用程序错误",
	-3:  "用户名或密码错误",
	-4:  "短信内容和备案的模板不一样",
	-5:  "签名不正确",
	-7:  "余额不足",
	-8:  "无可用通道或不在通道时间范围",
	-9:  "无效号码",
	-10: "签名内容不符合长度",
	-11: "用户有效期过期",
	-12: "短信平台-黑名单",
	-16: "接口请求过于频繁，余额接口 5s 秒一次，其他接口适当调整",
	-17: "非法 IP",
	-18: "Msg 内容格式错误",
	-19: "短信数量错误,小于 1 /大于 500(个性化)/大于 1000（群发）",
	-20: "号码错误或者黑名单",
	-23: "解密失败",
	-24: "短信包含用户敏感信息",
	-25: "用户被冻结",
	-26: "无效数据",
	-27: "请求参数错误",
	-28: "无效数据",
	-41: "指定短信模板类型错误或短信类型参数错误",
	-44: "自定义扩展号不符合规则（1-16 位数字）",
	-46: "用户黑名单",
	-47: "系统黑名单",
	-48: "号码超频拦截",
	-51: "超过设置的每月短信条数的限制",
	-54: "短信包含系统敏感信息",
}

func XMLToMap(xmlStr string) map[string]interface{} {

	params := make(map[string]interface{})
	// the current value stack
	values := make([]string, 0)
	decoder := xml.NewDecoder(strings.NewReader(xmlStr))

	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.CharData: // 标签内容
			values = append(values, string([]byte(token)))
		case xml.EndElement:
			if token.Name.Local == "xml" || token.Name.Local == "langs" {
				continue
			}
			params[token.Name.Local] = values[len(values)-1]
			// pop
			values = values[:]
		}
	}

	return params
}
