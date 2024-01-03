package smsConn

import (
	"encoding/json"
	"errors"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/liushuojia/go-sdk/utils"
)

const (
	smsAccessKeyID      = "LTAI5tJjjtY2QuySpotBJiUs"
	smsAaccessKeySecret = "oS9lbIJyGVEqYS7mz59Rrs89i6di4S"
	smsSignName         = "省心配"
	smsCodeTemplateCode = "SMS_222985010"
)

func New() *Conn {
	c := &Conn{}
	c.TemplateParam = make(map[string]string)
	return c
}

type Conn struct {
	Mobile        string            `json:"mobile"`         // 接受手机号码
	Type          string            `json:"type"`           // 类型 code 验证码， 其他待定
	TemplateParam map[string]string `json:"template_param"` // 短信模板变量对应的实际值，JSON格式。支持传入多个参数，示例：{"name":"张三","number":"15038****76"}。
}

func (obj *Conn) SetTemplateParam(key, value string) *Conn {
	obj.TemplateParam[key] = value
	return obj
}
func (obj *Conn) SetType(code string) *Conn {
	obj.Type = code
	return obj
}
func (obj *Conn) SetMobile(mobile string) *Conn {
	obj.Mobile = mobile
	return obj
}
func (obj *Conn) Send() error {
	switch obj.Type {
	case "code":
		if obj.Mobile == "" {
			return errors.New("参数传递错误")
		}
		if !utils.VerifyMobile(obj.Mobile) {
			return errors.New("手机格式错误")
		}
		if value, ok := obj.TemplateParam["code"]; !ok || value == "" {
			return errors.New("请传递验证码")
		}
		client, err := createClient(tea.String(smsAccessKeyID), tea.String(smsAaccessKeySecret))
		if err != nil {
			return err
		}
		sendSmsRequest := &dysmsapi20170525.SendSmsRequest{}
		sendSmsRequest.SetPhoneNumbers(obj.Mobile)
		sendSmsRequest.SetSignName(smsSignName)
		sendSmsRequest.SetTemplateCode(smsCodeTemplateCode)
		j, err := json.Marshal(obj.TemplateParam)
		if err != nil {
			return err
		}
		sendSmsRequest.SetTemplateParam(string(j))

		if _, err := client.SendSms(sendSmsRequest); err != nil {
			return err
		}
		return nil
	}
	return nil
}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func createClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}
