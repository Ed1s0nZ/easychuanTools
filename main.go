package main

import (
	"easychuanTools/tools"
	"flag"
	"fmt"
	"log"
	"strings"
)

//	func main() {
//		email := "testxx@test.com"
//		password := "dasdq2fva3q3f"
//		settext := "你好abc"
//		// 注册
//		regtoken, err := tools.SendRegRequest(email, password)
//		if err != nil {
//			fmt.Printf("Error: %v\n", err)
//		} else {
//			log.Println(regtoken)
//		}
//		//登录
//		logintoken, err := tools.SendLoginRequest(email, password)
//		if err != nil {
//			fmt.Printf("Error: %v\n", err)
//		} else {
//			fmt.Printf("Login successful! Token: %s\n", logintoken)
//		}
//		//发送文本
//		sentContenterr := tools.SendTextRequest(logintoken, settext)
//		if sentContenterr != nil {
//			fmt.Printf("Error: %v\n", sentContenterr)
//		} else {
//			fmt.Printf("Message sent successfully! Token: %s\n", logintoken)
//		}
//		//显示文本
//		text, GetContentRequesterr := tools.GetContentRequest(logintoken)
//		if GetContentRequesterr != nil {
//			fmt.Printf("Error: %v\n", GetContentRequesterr)
//		} else {
//			fmt.Printf("The text is: %s\n", text)
//		}
//	}
func main() {
	var (
		email    string
		password string
		settext  string
		regInfo  string
		gettext  string
	)

	flag.StringVar(&email, "u", "", "Email address")
	flag.StringVar(&password, "p", "", "Password")
	flag.StringVar(&settext, "s", "", "Text to send")
	flag.StringVar(&gettext, "g", "", "Get text")
	flag.StringVar(&regInfo, "r", "", "Register account (format: email/password)")
	flag.Parse()
	switch {
	case regInfo != "":
		splitRegInfo := strings.Split(regInfo, "/")
		if len(splitRegInfo) != 2 {
			fmt.Println("Invalid format for -r parameter. Please provide email and password separated by a slash.")
			return
		}
		email = splitRegInfo[0]
		password = splitRegInfo[1]

		regtoken, err := tools.SendRegRequest(email, password)
		if err != nil {
			fmt.Printf("Error registering account: %v\n", err)
			return
		}
		log.Println("Registration successful! Token:", regtoken)
		return // 注册完毕后直接结束程序
	case email != "" && password != "" && gettext != "":
		logintoken, err := tools.SendLoginRequest(email, password)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Login successful! Token: %s\n", logintoken)
		}
		text, GetContentRequesterr := tools.GetContentRequest(logintoken)
		if err != nil {
			fmt.Printf("Error: %v\n", GetContentRequesterr)
		} else {
			fmt.Printf("The text is: %s\n", text)
		}
	case email != "" && password != "" && settext != "":
		fmt.Println("settext", settext)
		logintoken, err := tools.SendLoginRequest(email, password)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Login successful! Token: %s\n", logintoken)
		}
		sentContenterr := tools.SendTextRequest(logintoken, settext)
		if sentContenterr != nil {
			fmt.Printf("Error: %v\n", sentContenterr)
		} else {
			fmt.Printf("Message sent successfully! Token: %s\n", logintoken)
		}
	default:
		if email == "" || password == "" || settext == "" {
			fmt.Println("Usage: ./easychuantools -u <email> -p <password> -s <text> -r <email/password>")
			return
		}

	
	}

}
