// Package errmail is my personal error-to-email package. You should make your
// own.
package errmail

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"strings"
	"sync/atomic"
)

const (
	envPassKey = "ERRMAIL_PASS"
)

var (
	gmailAddr = "smtp.gmail.com:587"
	addr      = gmailAddr
	from      = "cron@daaku.org"
	auth      = smtp.PlainAuth("", from, os.Getenv(envPassKey), "smtp.gmail.com")
	to        = []string{"n@daaku.org"}
	logged    int32
)

func init() {
}

func UseMailgun(user, pass string) {
	from = user
	auth = smtp.PlainAuth("", user, pass, "smtp.mailgun.org")
	addr = "smtp.mailgun.org:587"
}

// Send an error.
func Send(err error) {
	if atomic.CompareAndSwapInt32(&logged, 0, 1) {
		if os.Getenv(envPassKey) == "" && addr == gmailAddr {
			fmt.Fprintf(os.Stderr, "%s must be set for errmail\n", envPassKey)
		}
	}
	es := fmt.Sprintf("%+v", err)

	// subject upto the first newline
	subject := es
	if index := strings.Index(subject, "\n"); index != -1 {
		subject = subject[:index]
	}

	var b bytes.Buffer
	fmt.Fprintf(&b, "To: %s\r\n", strings.Join(to, ", "))
	fmt.Fprintf(&b, "Subject: %s\r\n", subject)
	fmt.Fprintf(&b, "\r\n%s", es)

	if e := smtp.SendMail(addr, auth, from, to, b.Bytes()); e != nil {
		fmt.Fprintln(os.Stderr, e)
	}
}
