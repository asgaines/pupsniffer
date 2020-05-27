package pupservice

import (
	"bytes"
	"net/smtp"
	"os"
)

func (p pupsvc) Mailman(buf *bytes.Buffer, recipients []string) error {
	subject := "Subject: New Pup(s) at the Boulder Humane Society!\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := buf.Bytes()

	msg := make([]byte, 0, len(subject)+len(mime)+len(body))

	msg = append(msg, []byte(subject)...)
	msg = append(msg, []byte(mime)...)
	msg = append(msg, body...)

	auth := smtp.PlainAuth(
		"",
		"andrew.s.gaines@gmail.com",
		os.Getenv("GMAIL_PASS"),
		"smtp.gmail.com",
	)

	sem := make(chan struct{}, 5)
	errs := make(chan error)

	go func() {
		for _, recipient := range recipients {
			sem <- struct{}{}
			go func(recipient string) {
				errs <- smtp.SendMail(
					"smtp.gmail.com:587",
					auth,
					"andrew.s.gaines@gmail.com",
					[]string{recipient},
					msg,
				)
				<-sem
			}(recipient)
		}
	}()

	for range recipients {
		if err := <-errs; err != nil {
			return err
		}
	}

	return nil
}
