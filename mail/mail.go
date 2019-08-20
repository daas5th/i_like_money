package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"reflect"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		send("hello there")
		return c.JSON(http.StatusOK, "send_OK")
	})

	e.Logger.Fatal(e.Start(":1331"))
}

func send(body string) {
	from, pass := ReadJson()
	to := "" //add reciever's mail address

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, success!")
}

type UserInfo struct {
	From string `json:"from"`
	Pass string `json:"pass"`
}

func ReadJson() (string, string) {
	jsonFile, err := os.Open("mail.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var userinfo UserInfo

	json.Unmarshal(byteValue, &userinfo)
	fmt.Println(reflect.TypeOf(userinfo.From))
	return userinfo.From, userinfo.Pass
}
