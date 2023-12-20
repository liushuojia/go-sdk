package dw

import (
	"github.com/fogleman/gg"
	"golang.org/x/image/draw"
	"image"
	"image/color"
	"io"
	"strings"
)

func New(height, width int) *Draw {
	return (&Draw{}).SetHeight(height).SetWidth(width).Context()
}

type Draw struct {
	Height int
	Width  int

	dc *gg.Context
}

func (d *Draw) SetHeight(height int) *Draw {
	d.Height = height
	return d
}
func (d *Draw) SetWidth(width int) *Draw {
	d.Width = width
	return d
}

func (d *Draw) Context() *Draw {
	d.dc = gg.NewContext(d.Width, d.Height)
	return d
}
func (d *Draw) Background(hexColor string) *Draw {
	d.dc.SetHexColor(hexColor)
	d.dc.Clear()
	return d
}
func (d *Draw) Rectangle(x, y, w, h float64) *Draw {
	d.dc.DrawRectangle(x, y, w, h)      // 添加一条以(50,50)为左上点、宽300、长233的矩形子路径
	d.dc.SetLineWidth(1)                // 设置画笔宽度为5像素
	d.dc.SetRGBA255(245, 34, 45, 255/2) // 设置画笔颜色为红色，且半透明
	d.dc.SetHexColor("#fafafa")         //
	d.dc.StrokePreserve()               // 使用当前颜色（红）描出当前路径（矩形），但不删除当前路径
	d.dc.SetHexColor("#fafafa")         // 设置画笔颜色为绿色
	d.dc.Fill()
	return d
}
func (d *Draw) Image(src string, x, y, w, h int, centerAlign bool) *Draw {
	if bg, err := gg.LoadImage(src); err == nil {
		bg = d.resizeImage(bg, w, h, true, 0, centerAlign)
		d.dc.DrawImage(bg, x, y) // 贴图（会受上述缩放倍率影响）
	}
	return d
}
func (d *Draw) Text(text string, x, y, w float64, align gg.Align) *Draw {
	strArr := d.changeText(text, w-20)
	if len(strArr) > 3 {
		strArr = strArr[0:3]
		strArr[2] = strArr[2] + "..."
	}
	d.dc.SetRGB255(0, 0, 0)
	d.dc.DrawStringWrapped(
		strings.Join(strArr, " "),
		x,
		y,
		0,
		0,
		w-20,
		1.5,
		align,
	)
	return d
}
func (d *Draw) Line(x, y, toX, toY float64, number string) *Draw {
	d.dc.SetLineWidth(2)
	d.dc.SetHexColor("#b25252")
	d.dc.DrawLine(x, y, toX, toY)
	d.dc.Stroke()

	d.dc.SetHexColor("#fafafa")
	d.dc.DrawCircle(toX, toY, 10) // 添加一条以(250,250)为圆心、200为半径的圆形子路径
	d.dc.Fill()

	d.dc.SetHexColor("#333333")
	d.dc.DrawStringWrapped(
		number,
		toX,
		toY,
		0.5,
		0.6,
		22,
		1.3,
		gg.AlignCenter,
	)

	//dc.SetRGB255(255, 222, 173)
	//dc.DrawCircle(float64(x), float64(y), 1) // 添加一条以(250,250)为圆心、200为半径的圆形子路径
	//dc.Fill()
	return d
}

func (d *Draw) Save(path string) error {
	return d.dc.SavePNG(path)
}
func (d *Draw) Writer(w io.Writer) error {
	/*
		srcFile, err := os.Create("./new_os.png")
		if err != nil {
			fmt.Println("open err", err)
			return
		}
		defer srcFile.Close()
		wr := bufio.NewWriter(srcFile)
		d.Writer(w)
	*/
	return d.dc.EncodePNG(w)
}

func (d *Draw) changeText(txt string, length float64) []string {
	txtArr := []rune(txt)
	var strArray []string
	str1 := ""
	for i := 0; i < len(txtArr); i++ {
		letter := string(txtArr[i : i+1])
		sWidth, _ := d.dc.MeasureString(str1 + letter)
		if sWidth >= length {
			strArray = append(strArray, str1)
			str1 = letter
		} else {
			str1 += letter
		}
	}
	strArray = append(strArray, str1)
	return strArray
}
func (d *Draw) resizeImage(img image.Image, width int, height int, keepRatio bool, fill int, centerAlign bool) image.Image {
	outImg := image.NewRGBA(image.Rect(0, 0, width, height))
	if !keepRatio {
		draw.BiLinear.Scale(outImg, outImg.Bounds(), img, img.Bounds(), draw.Over, nil)
		return outImg
	}

	if fill != 0 {
		fillColor := color.RGBA{R: uint8(fill), G: uint8(fill), B: uint8(fill), A: 255}
		draw.Draw(outImg, outImg.Bounds(), &image.Uniform{C: fillColor}, image.Point{}, draw.Src)
	}
	dst := d.calcResizedRect(width, img.Bounds(), height, centerAlign)
	draw.ApproxBiLinear.Scale(outImg, dst.Bounds(), img, img.Bounds(), draw.Over, nil)
	return outImg
}
func (d *Draw) calcResizedRect(width int, src image.Rectangle, height int, centerAlign bool) image.Rectangle {
	var dst image.Rectangle
	if width*src.Dy() < height*src.Dx() { // width/src.width < height/src.height
		ratio := float64(width) / float64(src.Dx())

		tH := int(float64(src.Dy()) * ratio)
		pad := 0
		if centerAlign {
			pad = (height - tH) / 2
		}
		dst = image.Rect(0, pad, width, pad+tH)
	} else {
		ratio := float64(height) / float64(src.Dy())
		tW := int(float64(src.Dx()) * ratio)
		pad := 0
		if centerAlign {
			pad = (width - tW) / 2
		}
		dst = image.Rect(pad, 0, pad+tW, height)
	}

	return dst
}
