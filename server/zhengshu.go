package main

//
//import (
//	"fmt"
//	"github.com/golang/freetype"
//	"github.com/golang/freetype/truetype"
//	"image"
//	"image/draw"
//	"image/png"
//	"io/ioutil"
//	"os"
//)
//
//var (
//	newTimes *truetype.Font // 字体
//	fontTtf  *truetype.Font // 字体
//)
//
//// 根据路径加载字体文件
//// path 字体的路径
//func loadFont(path string) (font *truetype.Font, err error) {
//	var fontBytes []byte
//	fontBytes, err = ioutil.ReadFile(path) // 读取字体文件
//	if err != nil {
//		err = fmt.Errorf("加载字体文件出错:%s", err.Error())
//		return
//	}
//	font, err = freetype.ParseFont(fontBytes) // 解析字体文件
//	if err != nil {
//		err = fmt.Errorf("解析字体文件出错,%s", err.Error())
//		return
//	}
//	return
//}
//func createHonoraryCertificate() (err error) {
//	// 根据路径打开模板文件
//	templateFile, err := os.Open("./image/certificate.png")
//	if err != nil {
//		return
//	}
//	defer templateFile.Close()
//	// 解码
//	templateFileImage, err := png.Decode(templateFile)
//	if err != nil {
//		return
//	}
//	// 新建一张和模板文件一样大小的画布
//	newTemplateImage := image.NewRGBA(templateFileImage.Bounds())
//	// 将模板图片画到新建的画布上
//	draw.Draw(newTemplateImage, templateFileImage.Bounds(), templateFileImage, templateFileImage.Bounds().Min, draw.Over)
//	// 获取图片的宽度和高度
//	width := newTemplateImage.Bounds().Dx()
//	height := newTemplateImage.Bounds().Dy()
//	fmt.Println(width, height)
//	// 加载字体文件  这里我们加载两种字体文件
//	newTimes, err = loadFont("./times.ttf")
//	if err != nil {
//		return
//	}
//	// 向图片中写入文字
//	// 在写入之前有一些准备工作
//	content := freetype.NewContext()
//	content.SetClip(newTemplateImage.Bounds())
//	content.SetDst(newTemplateImage)
//	content.SetSrc(image.Black) // 设置字体颜色
//	content.SetDPI(72)          // 设置字体分辨率
//	content.SetFont(newTimes)   // 设置字体样式，就是我们上面加载的字体
//
//	contentAuthors(content) // 写入作者信息
//	contentData(content)    // 写入数据信息
//	contentHash(content)    // 写入hash信息
//	contentDate(content)    // 写入日期信息
//
//	// 保存图片
//	dstFile, err := os.Create("./zhengshu.png")
//	if err != nil {
//		return
//	}
//	defer dstFile.Close()
//	png.Encode(dstFile, newTemplateImage)
//	fmt.Println("生成成功")
//	return
//}
//
//func contentAuthors(content *freetype.Context) {
//	// 一个字母大小占40宽，图片宽3508,高2427
//	authors := "author1 author1"
//	// 计算authors所占的宽度，然后计算出居中的x坐标
//	authorsWidth := len(authors) * 40
//	authorsX := (3508 - authorsWidth) / 2
//	content.DrawString(authors, freetype.Pt(authorsX, 1200))
//}
//func contentData(content *freetype.Context) {
//	content.SetFontSize(70) // 设置字体大小
//	data := "Certificate of acceptance for the manuscrip(2024 Mypapers International Conference on Blockchains) titled:Leveraging the power of internet of things and artificial intelligence in forest fire prevention, detection, and restoration: A comprehensive survey from MyPapers."
//	dataX := 400
//	dataY := 1300
//	for i := 0; i < len(data); i += 90 {
//		if i == 0 {
//			content.DrawString(data[i:i+80], freetype.Pt(dataX+110, dataY))
//			dataY += 80
//			i -= 10
//			continue
//		}
//		if i+90 > len(data) {
//			content.DrawString(data[i:], freetype.Pt(dataX, dataY))
//			break
//		}
//		content.DrawString(data[i:i+90], freetype.Pt(dataX, dataY))
//		dataY += 80
//	}
//}
//func contentHash(content *freetype.Context) {
//	content.SetFontSize(50) // 设置字体大小
//	transactionAddress := "0x16ce16c43bdb15e71fc74ded24f67111888ce708beffcbfb6dbdb412a8fefe3b"
//	content.DrawString(transactionAddress, freetype.Pt(1190, 1605))
//	blockAddress := "0x8ce4c18fd546719ccc8cd8d14ec3dc9224460efc872cc04035b3c9f9b2576d3a"
//	content.DrawString(blockAddress, freetype.Pt(1190, 1780))
//}
//func contentDate(content *freetype.Context) {
//	content.SetFontSize(70) // 设置字体大小
//	date := "2024.12.12"
//	content.DrawString(date, freetype.Pt(680, 1980))
//}
//func main() {
//	fmt.Println(createHonoraryCertificate())
//}
