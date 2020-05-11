package main

import (
	"github.com/lihaotian0607/qrcode"
	"log"
)

func main() {
}

func createForegroundImage() {
	qr, err := qrcode.New("http://www.baidu.com", qrcode.Highest)
	if err != nil {
		log.Fatalf("创建二维码失败 %s", err.Error())
	}

	qr.DisableBorder(true)

	qr.SetAvatar(&qrcode.Avatar{Src: "./static/avatar.jpg", Width: 60, Height: 60})
	qr.SetBackgroundImage(&qrcode.BackgroundImage{Src: "./static/bg.png", X: 70, Y: 55, Width: 270, Height: 270})
	qr.SetForegroundImage("./static/f.png")

	err = qr.WriteFile(400, "out.png")

	if err != nil {
		log.Fatalf("文件写入失败 %s", err.Error())
	}
}

func createBackgroundImage() {
	qr, err := qrcode.New("http://www.baidu.com", qrcode.Highest)
	if err != nil {
		log.Fatalf("创建二维码失败 %s", err.Error())
	}

	qr.DisableBorder(true)

	qr.SetAvatar(&qrcode.Avatar{Src: "./static/avatar.jpg", Width: 60, Height: 60})
	qr.SetBackgroundImage(&qrcode.BackgroundImage{Src: "./static/bg.png", X: 70, Y: 55, Width: 270, Height: 270})

	err = qr.WriteFile(400, "out.png")

	if err != nil {
		log.Fatalf("文件写入失败 %s", err.Error())
	}
}

func createAvatar() {
	qr, err := qrcode.New("http://www.baidu.com", qrcode.Highest)
	if err != nil {
		log.Fatalf("创建二维码失败 %s", err.Error())
	}

	qr.DisableBorder(true)

	qr.SetAvatar(&qrcode.Avatar{Src: "./static/avatar.jpg", Width: 60, Height: 60})

	err = qr.WriteFile(400, "avatar.png")

	if err != nil {
		log.Fatalf("文件写入失败 %s", err.Error())
	}
}
