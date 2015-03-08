package email

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type AnoMail struct {
	Subject         string   //主题
	ReceiverPostfix string   //默认的收件人邮箱后缀，没有添加
	MailType        string   //邮件类型
	From            string   //发件人
	To              []string //收件人
	Cc              []string //抄送
	Bcc             []string //密送
}

//结构体创建
func New(subject string, mail_from string, mail_tos []string) *AnoMail {
	Mail := &AnoMail{
		Subject:         subject,
		ReceiverPostfix: "@xiaomi.com",
		MailType:        "text/html",
		From:            mail_from,
		To:              mail_tos,
		Cc:              []string{},
		Bcc:             []string{},
	}

	return Mail
}

//发送邮件
func (this AnoMail) Send(content string) error {
	head_string := this.__make_head()
	mail_string := head_string + content

	//创建一个临时文件
	f, _ := ioutil.TempFile("/tmp", "NornsMail")
	temp_name := f.Name()
	f.WriteString(mail_string)
	f.Close()

	dedump_map := make(map[string]bool, len(this.To)+len(this.Cc)+len(this.Bcc))
	all_receiver := append(this.To, this.Cc...)
	all_receiver = append(all_receiver, this.Bcc...)
	for _, user_mail := range all_receiver {
		if _, ok := dedump_map[user_mail]; ok {
			continue
		}
		dedump_map[user_mail] = true
		__execute(temp_name, user_mail)
	}
	return nil

}

func __execute(temp_file string, receiver string) error {
	cmd := exec.Command("/usr/sbin/sendmail", receiver)
	fmt.Println(temp_file)
	file_handler, o_err := os.Open(temp_file)
	if o_err != nil {
		return o_err
	}
	cmd.Stdin = file_handler
	result, ex_err := cmd.Output()
	fmt.Println(result, ex_err)
	if ex_err != nil {
		return ex_err
	}

	return nil
}

//可以修改默认的收件人后缀，没有加后缀的收件人，将添加此后缀
//默认@xiaomi.com
func (this *AnoMail) SetReceiverPostfix(postfix string) error {
	this.ReceiverPostfix = postfix
	return nil
}

//用来添加Cc的人
func (this *AnoMail) SetCc(cc []string) error {
	this.Cc = cc
	return nil
}

//用来添加密送的人
func (this *AnoMail) SetBcc(bcc []string) error {
	this.Bcc = bcc
	return nil
}

//可以修改邮件的类型
//支持类型：
//text/plain 纯文本
//text/html  HTML文档
//text/xml   XML文档
func (this *AnoMail) SetType(mail_type string) error {
	types := map[string]bool{
		"text/plain": true,
		"text/html":  true,
		"text/xml":   true,
	}
	_, ok := types[mail_type]
	if true == ok {
		this.MailType = mail_type
		return nil
	} else {
		return errors.New("cannot support your type ...")
	}
}

//根据结构体内容，整合邮件头部
func (this AnoMail) __make_head() string {
	subject := fmt.Sprintf("Subject : %s\n", this.Subject)
	mail_type := fmt.Sprintf("Content-Type : %s\n", this.MailType)
	mail_from := fmt.Sprintf("From : %s\n", this.__make_addr_for_body([]string{this.From}))
	mail_to := fmt.Sprintf("To : %s\n", this.__make_addr_for_body(this.To))
	mail_cc := fmt.Sprintf("Cc : %s\n", this.__make_addr_for_body(this.Cc))
	ret := subject + mail_type + mail_from + mail_to + mail_cc + "\n"
	return ret
}

func (this AnoMail) __make_addr_for_body(mails []string) string {
	mails = __add_mail_postfix(mails, this.ReceiverPostfix)
	var temp_list []string
	for _, mail := range mails {
		temp_list = append(temp_list, "<"+mail+">")
	}
	addrs := strings.Join(temp_list, ";")
	return addrs
}

//检查邮件是否有后缀，没有记得加上
func __add_mail_postfix(receivers []string, postfix string) []string {
	for index, name := range receivers {
		if true == strings.Contains(name, "@") {
			continue
		} else {
			receivers[index] = name + postfix
		}
	}
	return receivers
}
