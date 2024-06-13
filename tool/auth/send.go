package auth

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

var (
	Password  = "cldhjlxvduxdbihc"
	Subject   = "验证码"
	SendEmail = "368809634@qq.com"
	prefix    = "【tongchi_mall】验证码:"
	suffix    = "用于邮箱身份验证，5分钟内有效，请勿泄露和转发。如非本人操作，请忽略此短信。"
)

func Send(from, Password, to, Subject, text string) {

	//from：发送方邮箱
	//password：授权码
	//to：接受方邮箱
	m := gomail.NewMessage()

	//发送人
	m.SetHeader("From", from)
	//接收人
	m.SetHeader("To", to)
	//抄送人
	//m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
	//主题
	m.SetHeader("Subject", Subject)
	//内容
	text = prefix + text + suffix
	m.SetBody("text/plain", text)
	//附件
	//m.Attach("./myIpPic.png")

	//拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer("smtp.qq.com", 587, SendEmail, Password)

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("DialAndSend err %v:", err)
		panic(err)
	}
	fmt.Printf("send mail success\n")

}
