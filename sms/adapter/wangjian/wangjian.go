package wangjian

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/x-io/gen.module/sms/core"
)

//https://utf8api.smschinese.cn/?Uid=本站用户名&Key=接口安全秘钥&smsMob=手机号码&smsText=验证码:8888

type engine struct {
	core.Engine
}

func New(c core.Engine) *engine {
	return &engine{
		c,
	}
}

// Send sendSms
func (w *engine) Send(phone string, text string) (int, error) {
	if w.Uid == "" {
		return 1, nil
	}

	params := "?Uid=${uid}&Key=${key}&smsMob=${phone}&smsText=${text}"

	u := strings.Replace(params, "${phone}", url.QueryEscape(phone), -1)
	u = strings.Replace(u, "${text}", url.QueryEscape(text), -1)
	u = strings.Replace(u, "${uid}", url.QueryEscape(w.Uid), -1)
	u = strings.Replace(u, "${key}", url.QueryEscape(w.Key), -1)

	resp, err := http.Get(w.Domain + u)
	if err != nil {
		return -100, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return -100, err
	}

	code, err := strconv.Atoi(string(respBody))
	if err != nil {
		return -100, err
	}
	if code > 0 {
		return code, nil
	}

	switch code {
	case -1:
		return -1, fmt.Errorf("没有该用户账户")
	case -2:
		return -2, fmt.Errorf("接口密钥不正确")
	case -3:
		return -3, fmt.Errorf("短信数量不足")
	case -4:
		return -4, fmt.Errorf("手机号格式不正确")
	case -6:
		return -6, fmt.Errorf("IP限制")
	case -11:
		return -11, fmt.Errorf("该用户被禁用")
	case -14:
		return -14, fmt.Errorf("短信内容出现非法字符")
	case -21:
		return -21, fmt.Errorf("MD5接口密钥加密不正确")
	case -41:
		return -41, fmt.Errorf("手机号码为空")
	case -42:
		return -42, fmt.Errorf("短信内容为空")
	case -51:
		return -51, fmt.Errorf("短信签名格式不正确")
	case -52:
		return -52, fmt.Errorf("短信签名太长 建议签名10个字符以内")
	default:
		return -100, fmt.Errorf("未知错误")
	}
}
