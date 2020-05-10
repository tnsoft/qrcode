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

	qr.SetBackground(&qrcode.Background{
		Src:    "../static/3.png",
		X:      70,
		Y:      55,
		Width:  270,
		Height: 270,
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
