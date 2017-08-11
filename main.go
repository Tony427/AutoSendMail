package main

import (
	"crypto/tls"

	"github.com/go-toast/toast"
	"gopkg.in/gomail.v2"
)

func main() {
	notification("Meaningless email Start!!")
	m := gomail.NewMessage()
	m.SetHeader("From", "u@pic.net.tw")

	m.SetHeader("To", "aaa@abc.com")

	m.SetHeader("Cc", "bbb@abc.com", "ccc@abc.com", "ddd@abc.com", "eee@abc.com")

	m.SetHeader("Subject", "Hola cómo estás")
	m.SetBody("text/html", `<h1>mi nombre es Tony</h1>`)

	d := gomail.NewDialer("smtp.office365.com", 587, "urmail@pic.net.tw", "urpassword")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := d.DialAndSend(m)

	if err != nil {
		notification("Oh....No....somethiing went wrong Q^Q")
		panic(err)
	} else {
		notification("Meaningless email success !!")
	}

	// Send emails using d.
}

func notification(message string) error {
	notification := toast.Notification{
		AppID:   "App name",
		Title:   "notify title",
		Message: message,
		//Icon:    "go.png", // This file must exist (remove this line if it doesn't)
		Actions: []toast.Action{
			{Type: "protocol", Label: "I'm a button", Arguments: ""},
			{Type: "protocol", Label: "Me too!", Arguments: ""},
		},
	}
	err := notification.Push()
	if err != nil {
		panic(err)
	}
	return nil
}
