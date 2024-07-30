package main

import (
	"easychuanApi/tools"
	"fmt"
	"log"
)

func main() {
	email := "testxx@xx.com"
	password := "xxx123qwe."
	settext := "你好abc"
	// 注册
	regtoken, err := tools.SendRegRequest(email, password)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		log.Println(regtoken)
	}
	//登录
	logintoken, err := tools.SendLoginRequest(email, password)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Login successful! Token: %s\n", logintoken)
	}
	//发送文本
	sentContenterr := tools.SendTextRequest(logintoken, settext)
	if err != nil {
		fmt.Printf("Error: %v\n", sentContenterr)
	} else {
		fmt.Printf("Message sent successfully! Token: %s\n", logintoken)
	}
	//展示文本
	text, GetContentRequesterr := tools.GetContentRequest(logintoken)
	if err != nil {
		fmt.Printf("Error: %v\n", GetContentRequesterr)
	} else {
		fmt.Printf("The text is: %s\n", text)
	}
}
