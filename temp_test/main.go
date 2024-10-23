package main

import "net/smtp"

func main() {

	from := "ubelwertyas@gmail.com"
	password := "nibe ipyf shxq wtzt"

	toEmailAddress := "almi.body@mail.ru"
	to := []string{toEmailAddress}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	subject := "Subject: This is the subject of the mail\n"
	body := "This is the body of the mail"
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		panic(err)
	}
}
