package email

import (
	_ "fmt"
)

func SendSimple(subject string, mail_from string, mail_to []string, content string) error {
	Mail := New(subject, mail_from, mail_to)
	err := Mail.Send(content)
	return err
}

func SendWithCc(subject string, mail_from string, mail_to []string, mail_cc []string, content string) error {
	Mail := New(subject, mail_from, mail_to)
	Mail.SetCc(mail_cc)
	err := Mail.Send(content)
	return err
}

func SendWithBcc(subject string, mail_from string, mail_to []string, mail_cc []string, mail_bcc []string, content string) error {
	Mail := New(subject, mail_from, mail_to)
	Mail.SetCc(mail_cc)
	Mail.SetBcc(mail_bcc)
	err := Mail.Send(content)
	return err
}
