package qrcode

import (
	"fmt"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
	"image"
	"image/color"
	"image/draw"
	"os"

	_ "image/jpeg"
)

type Background struct {
	Src    string
	X      int
	Y      int
	Width  int
	Height int
}

type Foreground struct {
	Src string
}

type Avatar struct {
	Src    string // 头像地址
	Width  int    // 头像宽度
	Height int    // 头像高度
	Round  int    // 圆角
}

type QrCode struct {
	*qrcode.QRCode
	Avatar     *Avatar
	Foreground *Foreground
	Background *Background
}

func New(content string, level qrcode.RecoveryLevel) (*QrCode, error) {

	qr, err := qrcode.New(content, level)
	if err != nil {
		return nil, err
	}

	return &QrCode{QRCode: qr}, nil
}

const (
	// Level L: 7% error recovery.
	Low qrcode.RecoveryLevel = iota

	// Level M: 15% error recovery. Good default choice.
	Medium

	// Level Q: 25% error recovery.
	High

	// Level H: 30% error recovery.
	Highest
)

func (q *QrCode) SetAvatar(avatar *Avatar) {
	q.Avatar = avatar
}

func (q *QrCode) SetForeground(src string) {
	q.Foreground = &Foreground{Src: src}
}

func (q *QrCode) SetBackground(bg *Background) {
	q.Background = bg
}

func (q *QrCode) Generate(size int) (image.Image, error) {
	img := q.Image(size)
	var err error

	if q.Foreground != nil {
		img, err = q.GenerateForeground(img)
		if err != nil {
			return nil, err
		}
	}

	if q.Avatar != nil {
		img, err = q.GenerateAvatar(img)
		if err != nil {
			return nil, err
		}
	}

	if q.Background != nil {
		img, err = q.GenerateBackground(img)
		if err != nil {
			return nil, err
		}
	}

	return img, nil
}

func (q *QrCode) GenerateAvatar(img image.Image) (image.Image, error) {
	avatar, err := os.Open(q.Avatar.Src)
	if err != nil {
		return nil, fmt.Errorf("打开头像文件失败 %s", err.Error())
	}

	defer avatar.Close()

	decode, _, err := image.Decode(avatar)
	if err != nil {
		return nil, err
	}

	decode = resize.Resize(uint(q.Avatar.Width), uint(q.Avatar.Height), decode, resize.Lanczos3)

	b := img.Bounds()

	// 设置为居中
	offset := image.Pt((b.Max.X-decode.Bounds().Max.X)/2, (b.Max.Y-decode.Bounds().Max.Y)/2)

	m := image.NewRGBA(b)

	draw.Draw(m, b, img, image.Point{X: 0, Y: 0}, draw.Src)

	draw.Draw(m, decode.Bounds().Add(offset), decode, image.Point{X: 0, Y: 0}, draw.Over)

	return m, err
}

func (q *QrCode) GenerateForeground(img image.Image) (image.Image, error) {
	file, err := os.Open(q.Foreground.Src)
	if err != nil {
		return nil, fmt.Errorf("打开前景图文件失败 %s", err.Error())
	}

	defer file.Close()

	decode, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	// 获取二维码的宽高
	width, height := img.Bounds().Max.X, img.Bounds().Max.Y

	foregroundW, foregroundH := decode.Bounds().Max.X, decode.Bounds().Max.Y

	if width != foregroundW || height != foregroundH {
		// 如果不一致将填充图剪裁
		decode = resize.Resize(uint(width), uint(height), decode, resize.Lanczos3)
	}

	b := img.Bounds()
	imgSet := image.NewRGBA(b)

	for y := 0; y < img.Bounds().Max.X; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {

			qrImgColor := img.At(x, y)

			// 检测图片颜色 如果rgb值是 255 255 255 255 则像素点为白色 跳过
			// 如果rgba值是 0 0 0 0 则为透明色 跳过

			c := qrImgColor.(color.Gray16)

			if c == color.White {
				continue
			}

			// 获取要填充的图片的颜色
			bgImgColor := decode.At(x, y)

			// 填充颜色
			switch decode.(type) {
			case *image.RGBA:
				c := bgImgColor.(color.RGBA)
				imgSet.Set(x, y, color.RGBA{R: c.R, G: c.G, B: c.B, A: c.A})
			case *image.YCbCr:
				c := bgImgColor.(color.YCbCr)
				img.(draw.Image).Set(x, y, color.YCbCr{Y: c.Y, Cb: c.Cb, Cr: c.Cr})

			default:
				fmt.Println("error")
			}

		}
	}

	return imgSet, nil
}

func (q *QrCode) GenerateBackground(img image.Image) (image.Image, error) {
	file, err := os.Open(q.Background.Src)
	if err != nil {
		return nil, fmt.Errorf("打开背景图文件失败 %s", err.Error())
	}

	img = resize.Resize(uint(q.Background.Width), uint(q.Background.Height), img, resize.Lanczos3)

	defer file.Close()

	bg, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	offset := image.Pt(q.Background.X, q.Background.Y)

	b := bg.Bounds()

	m := image.NewRGBA(b)

	draw.Draw(m, b, bg, image.Point{X: 0, Y: 0}, draw.Src)

	draw.Draw(m, img.Bounds().Add(offset), img, image.Point{X: 0, Y: 0}, draw.Over)

	return m, nil
}
