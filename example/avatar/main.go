package main

import (
	"fmt"
	"github.com/lihaotian0607/qrcode"
	"image/png"
	"log"
	"os"
)

func main() {
	var qr, err = qrcode.New("https://www.xxx.com/", qrcode.Highest)
	if err != nil {
		fmt.Println("创建二维码失败")
		return
	}

	qr.QRCode.DisableBorder = true

	qr.SetAvatar(&qrcode.Avatar{
		Src:    "../static/1.jpg",
		Width:  60,
		Height: 60,
		Round:  10,
	})

	img, err := qr.Generate(300)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	file, err := os.Create("out.png")

	if err != nil {
		log.Fatalf("创建文件失败 %s", err.Error())
		return
	}

	err = png.Encode(file, img)
	if err != nil {
		log.Fatalf("文件写入失败 %s", err.Error())
	}
}
