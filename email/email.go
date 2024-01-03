package email

import (
	"errors"
	"fmt"
	"gopkg.in/gomail.v2"
	"mime"
	"regexp"
	"strings"
)

func New() *Email {
	return &Email{}
}

type Email struct {
	Subject     string       `json:"subject"`
	Body        string       `json:"body"`
	From        string       `json:"from"`
	CompanyName string       `json:"company_name"`
	MailTo      []Mail       `json:"mail_to"`
	MailCc      []Mail       `json:"mail_cc"`
	MailBcc     []Mail       `json:"mail_bcc"`
	MailAttach  []MailAttach `json:"mail_attach"`
	Host        MailHost     `json:"host"`
}
type Mail struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
type MailAttach struct {
	Name string `json:"name"`
	Path string `json:"path"`
}
type MailHost struct {
	Account string `json:"account"`
	Name    string `json:"name"`
	Passwd  string `json:"passwd"`
	Smtp    string `json:"smtp"`
	Port    int    `json:"port"`
}

func (obj *Email) SetCompanyName(companyName string) *Email {
	obj.CompanyName = companyName
	return obj
}
func (obj *Email) SetHost(account, name, passwd, smtp string, port int) *Email {
	obj.Host = MailHost{
		Account: account,
		Name:    name,
		Passwd:  passwd,
		Smtp:    smtp,
		Port:    port,
	}
	return obj
}
func (obj *Email) SetSubject(subject string) *Email {
	obj.Subject = subject
	return obj
}
func (obj *Email) SetBody(body string) *Email {
	obj.Body = body
	return obj
}
func (obj *Email) SetFrom(from string) *Email {
	obj.From = from
	return obj
}

func (obj *Email) AddMailTo(name, address string) *Email {
	obj.MailTo = append(obj.MailTo, Mail{Name: name, Address: address})
	return obj
}
func (obj *Email) AddMailCc(name, address string) *Email {
	obj.MailCc = append(obj.MailCc, Mail{Name: name, Address: address})
	return obj
}
func (obj *Email) AddMailBcc(name, address string) *Email {
	obj.MailBcc = append(obj.MailBcc, Mail{Name: name, Address: address})
	return obj
}
func (obj *Email) AddMailAttach(name, path string) *Email {
	obj.MailAttach = append(obj.MailAttach, MailAttach{Name: name, Path: path})
	return obj
}

func (obj *Email) VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	email = strings.ToLower(email)
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
func (obj *Email) Send() error {
	if len(obj.MailTo) == 0 {
		return errors.New("收件人为空")
	}

	for _, v := range obj.MailTo {
		if !obj.VerifyEmailFormat(v.Address) {
			return errors.New("收件人中有错误邮件地址")
		}
	}

	m := gomail.NewMessage(
		gomail.SetEncoding(gomail.Base64),
	)

	//这种方式可以添加别名，即“XX官方”
	m.SetHeader("From", m.FormatAddress(obj.Host.Account, obj.Host.Name))
	// 说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	//m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“FB Sample”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	//m.SetHeader("From", mailConn["user"])

	var toArray []string
	for _, v := range obj.MailTo {
		toArray = append(toArray, m.FormatAddress(v.Address, v.Name))
	}
	m.SetHeader("To", toArray...) //发送给多个用户

	var ccArray []string
	for _, v := range obj.MailCc {
		ccArray = append(ccArray, m.FormatAddress(v.Address, v.Name))
	}
	if len(ccArray) > 0 {
		m.SetHeader("Cc", ccArray...) // 抄送
	}

	var bccArray []string
	for _, v := range obj.MailBcc {
		bccArray = append(bccArray, m.FormatAddress(v.Address, v.Name))
	}
	if len(ccArray) > 0 {
		m.SetHeader("Bcc", bccArray...) // 暗送
	}

	m.SetHeader("Subject", obj.Subject) //设置邮件主题
	m.SetBody("text/html", obj.Body)    //设置邮件正文

	if len(obj.MailAttach) > 0 {
		for _, attach := range obj.MailAttach {
			if attach.Name == "" || attach.Path == "" {
				continue
			}
			m.Attach(attach.Path,
				gomail.Rename(attach.Name),
				gomail.SetHeader(map[string][]string{
					"Content-Disposition": {
						fmt.Sprintf(`attachment; filename="%s"`, mime.BEncoding.Encode("UTF-8", attach.Name)),
					},
				}),
			)
		}
	}

	return gomail.
		NewDialer(
			obj.Host.Smtp,
			obj.Host.Port,
			obj.Host.Account,
			obj.Host.Passwd,
		).DialAndSend(m)
}
