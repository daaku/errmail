// Package errmail is a way to format a error into an email.
package errmail

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

type Logger struct {
	Logger *log.Logger
	Addr   string
	Auth   smtp.Auth
	From   string
	To     []string
}

func (l *Logger) Log(err error) {
	es := err.Error()

	// subject upto the first newline
	subject := es
	if index := strings.Index(subject, "\n"); index != -1 {
		subject = subject[:index]
	}

	var b bytes.Buffer
	fmt.Fprintf(&b, "To: %s\r\n", strings.Join(l.To, ", "))
	fmt.Fprintf(&b, "Subject: %s\r\n", subject)
	fmt.Fprintf(&b, "\r\n%s", es)

	if e := smtp.SendMail(l.Addr, l.Auth, l.From, l.To, b.Bytes()); e != nil {
		if l.Logger == nil {
			log.Println(e)
		} else {
			l.Logger.Println(e)
		}
	}
}
