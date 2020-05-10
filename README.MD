## 设置头像
```golang
    qr.SetAvatar(&qrcode.Avatar{
		Src:    "../static/1.jpg",
		Width:  60,
		Height: 60,
		Round:  10,
	})
```

## 设置背景图
```
qr.SetBackground(&qrcode.Background{
		Src:    "../static/3.png",
		X:      70,
		Y:      55,
		Width:  270,
		Height: 270,
	})
```

## 设置前景图
```
qr.SetForeground("../static/2.png")
```

![image](https://github.com/lihaotian0607/qrcode/blob/master/resources/20200510170738.png)
![image](https://github.com/lihaotian0607/qrcode/blob/master/resources/20200510170741.png)
![image](https://github.com/lihaotian0607/qrcode/blob/master/resources/20200510170755.png)
![image](https://github.com/lihaotian0607/qrcode/blob/master/resources/20200510170802.png)
![image](https://github.com/lihaotian0607/qrcode/blob/master/resources/20200510170810.png)
![image](https://github.com/lihaotian0607/qrcode/blob/master/resources/20200510170813.png)





