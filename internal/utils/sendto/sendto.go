package sendto

import (
	"fmt"
	"strings"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\"\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ","))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}

func SendTextEmail(to []string, from string, otp string) error {
	// ### TODO: ###

	// contentEmail := Mail{
	// 	From:    EmailAddress{Address: from, Name: "test"},
	// 	To:      to,
	// 	Subject: "OTP Verification",
	// 	Body:    fmt.Sprintf("Your OTP is %s. Please enter it to verify your account.", otp),
	// }
	// ....
	// messageMail := BuildMessage(contentEmail)

	//send smtp
	// auth := smtp.PlainAuth("")
	return nil
}
