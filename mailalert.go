package main

import (
  "strconv"
  "crypto/tls"
  "gopkg.in/gomail.v2"
)

const temphigh = 30.0
const humhigh = 30.0

const mailfrom = "weather@mail.l4rs.net"
const mailto = "me@l4rs.net"
const mailserver = "mail.l4rs.net"
const mailport = 587
const mailpassword = "PASSWORD"

// MailAlert checks if the temperature and humidity values are greater or equal
// to the hard coded threshold constants (temphigh and humhigh), if yes an email
// will be sent to the configured constant (mailto). Temperature and humidity
// will alert in seperate emails.
// The constants mailfrom (will be used as username too), mailto, mailserver,
// mailport and mailpassword need to be set.
func MailAlert (temp string, hum string) {
  tempnew, err := strconv.ParseFloat(temp, 64)
  if err != nil {
  // insert error handling here
  }
  humnew, err := strconv.ParseFloat(hum, 64)
  if err != nil {
  // insert error handling here
  }

  if tempnew >= temphigh {
    m := gomail.NewMessage()
  	m.SetHeader("From", mailfrom)
  	m.SetHeader("To", mailto)
  	m.SetHeader("Subject", "Temp to high!")
  	m.SetBody("text/html", "Temp is equal or above threshold")

    d := gomail.NewDialer(mailserver, mailport, mailfrom, mailpassword)
    d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

  	if err := d.DialAndSend(m); err != nil {
  		panic(err)
  	}
}

if humnew >= humhigh {
  m := gomail.NewMessage()
  m.SetHeader("From", mailfrom)
  m.SetHeader("To", mailto)
  m.SetHeader("Subject", "Hum to high!")
  m.SetBody("text/html", "Hum is equal or above threshold")

  d := gomail.NewDialer(mailserver, mailport, mailfrom, mailpassword)
  d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

  if err := d.DialAndSend(m); err != nil {
    panic(err)
  }
}
}
