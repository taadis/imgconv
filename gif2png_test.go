package imgconv

import (
	"bytes"
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"testing"
)

func TestGIFToPNGConverter_Convert(t *testing.T) {
	// 创建一个简单的GIF图像用于测试
	gifBuffer := new(bytes.Buffer)
	// 创建一个基本的调色板
	palette := []color.Color{color.White, color.Black}
	img := gif.GIF{
		Image: []*image.Paletted{image.NewPaletted(image.Rect(0, 0, 2, 2), palette)},
		Delay: []int{0},
	}
	if err := gif.EncodeAll(gifBuffer, &img); err != nil {
		t.Fatalf("创建测试GIF失败: %v", err)
	}

	// 创建转换器
	converter := NewGIFToPNGConverter()

	// 准备输出buffer
	outBuffer := new(bytes.Buffer)

	// 执行转换
	if err := converter.Convert(gifBuffer, outBuffer); err != nil {
		t.Fatalf("转换失败: %v", err)
	}

	// 验证输出是否为有效的PNG
	_, err := png.Decode(bytes.NewReader(outBuffer.Bytes()))
	if err != nil {
		t.Fatalf("输出不是有效的PNG: %v", err)
	}
}
