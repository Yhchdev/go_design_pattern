package main

import "fmt"

type Computer interface {
	InsertIntoLightningPort()
}

type Client struct {
}

func (t Client) InsertLightIntoComputer(c Computer) {
	c.InsertIntoLightningPort()
}

type Mac struct {
}

func (m Mac) InsertIntoLightningPort() {
	fmt.Println("给mac电脑插入雷电接口")
}

type Windows struct {
}

func (m Windows) InsertIntoUsbPort() {
	fmt.Println("给windows电脑插入usb接口")
}

type WindowsAdapter struct {
	windows Windows
}

func (w WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("转换雷电接口为usb接口")
	w.windows.InsertIntoUsbPort()
}

func main() {
	mac := Mac{}
	client := Client{}

	client.InsertLightIntoComputer(mac)

	windows := Windows{}

	adapter := WindowsAdapter{windows: windows}

	client.InsertLightIntoComputer(adapter)
}


// output
// 给mac电脑插入雷电接口
// 转换雷电接口为usb接口
// 给windows电脑插入usb接口
