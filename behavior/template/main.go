package main

import "fmt"

type IVerificationCode interface {
	genCode() string
	saveCode(code string) error
	getMsg(code string) string
	sendMsg(msg string) error
}

func genAndSendCode(opt IVerificationCode) error {
	code := opt.genCode()

	if err := opt.saveCode(code); err != nil {
		return err
	}
	msg := opt.getMsg(code)

	if err := opt.sendMsg(msg); err != nil {
		return err
	}

	return nil
}

// 公共部分抽象
type CodePart struct {
}

func (c CodePart) genCode() string {
	return "1234"
}

func (c CodePart) getMsg(code string) string {
	return fmt.Sprintf("你的验证码是：%s", code)

}

func (c CodePart) saveCode(code string) error {
	fmt.Printf("服务端保存了验证码：%s\n", code)
	return nil
}

type SmsCode struct {
	CodePart
}

func (s SmsCode) sendMsg(code string) error {
	fmt.Println("通过短信的方式发送了验证码")
	return nil
}

type EmailCode struct {
	CodePart
}

func (e EmailCode) sendMsg(code string) error {
	fmt.Println("通过邮件的方式发送了验证码")
	return nil
}

func main() {

	sms := SmsCode{}

	if err := genAndSendCode(sms); err != nil {
		return
	}

	email := EmailCode{}

	if err := genAndSendCode(email); err != nil {
		return
	}

}
