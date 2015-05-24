// Package errmail is my personal error-to-email package. You should make your
// own.
package errmail

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
)

const from = "cron@daaku.org"

var (
	auth = smtp.PlainAuth("", from, os.Getenv("ERRMAIL_PASS"), "smtp.gmail.com")
	to   = []string{"n@daaku.org"}
)

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
		log.Println(e)
	}
}
