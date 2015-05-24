// Package errmail is my personal error-to-email package. You should make your
// own.
package errmail

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

const (
	envPassKey = "ERRMAIL_PASS"
	from       = "cron@daaku.org"
)

var (
	auth = smtp.PlainAuth("", from, os.Getenv(envPassKey), "smtp.gmail.com")
	to   = []string{"n@daaku.org"}
)

func init() {
	if os.Getenv(envPassKey) == "" {
		fmt.Fprintf(os.Stderr, "%s must be set for errmail\n", envPassKey)
	}
}

// Send an error.
func Send(err error) {
	es := err.Error()

	// subject upto the first newline
	subject := es
	if index := strings.Index(subject, "\n"); index != -1 {
		subject = subject[:index]
	}

	var b bytes.Buffer
	fmt.Fprintf(&b, "To: %s\r\n", strings.Join(to, ", "))
	fmt.Fprintf(&b, "Subject: %s\r\n", subject)
	fmt.Fprintf(&b, "\r\n%s", es)

	if e := smtp.SendMail("smtp.gmail.com:587", auth, from, to, b.Bytes()); e != nil {
		fmt.Fprintln(os.Stderr, e)
	}
}
