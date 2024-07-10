package pdf

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"io"
	"strings"
)

func dealImageByBitPerixel(reader io.Reader, width int, height int, bitPerPixel int, ColorSpace string) (out io.Reader) {
	// config := color.RGBA64Model

	// 创建一个新的图像
	img := image.NewRGBA64(image.Rect(0, 0, width, height))
	buf, _ := io.ReadAll(reader)
	// 填充图像数据，这里只是一个示例，具体的数据处理逻辑需要根据你的需求来编写
	ColorSpace = strings.TrimPrefix(ColorSpace, "/")
	switch ColorSpace {
	case "DeviceGray":

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				offset := y*width + x
				g := buf[offset]
				img.Set(x, y, color.Gray{Y: g})
			}
		}
	case "DeviceRGB":
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				offset := (y*width + x) * 3
				r := buf[offset+0]
				g := buf[offset+1]
				b := buf[offset+2]
				a := uint8(255)
				// a := buf[(y*width+x)*4+3] // 如果没有alpha通道，这里可以是255
				img.Set(x, y, color.RGBA{R: r, G: g, B: b, A: a})
				// img.SetRGBA64(x, y, color.RGBA64{R: grayValue, G: grayValue, B: grayValue, A: 65535})
			}
		}

	}

	buffer := bytes.NewBuffer([]byte{})
	png.Encode(buffer, img)
	return buffer
}
