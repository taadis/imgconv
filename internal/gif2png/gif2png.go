package gif2png

import (
	"image/gif"
	"image/png"
	"io"
)

// GIFToPNGConverter 实现了Converter接口，用于GIF到PNG的转换
type GIFToPNGConverter struct{}

// NewGIFToPNGConverter 创建一个新的GIF到PNG转换器
func NewGIFToPNGConverter() *GIFToPNGConverter {
	return &GIFToPNGConverter{}
}

// Convert 实现Converter接口，将GIF图片转换为PNG格式
func (c *GIFToPNGConverter) Convert(input io.Reader, output io.Writer) error {
	// 解码 GIF 图像
	gifImage, err := gif.DecodeAll(input)
	if err != nil {
		return err
	}

	// 如果 GIF 包含多帧，这里只处理第一帧
	// 如果需要处理所有帧，可以选择保存为多张 PNG 或其他方式
	pngImage := gifImage.Image[0]

	// 编码为 PNG 格式
	err = png.Encode(output, pngImage)
	if err != nil {
		return err
	}

	return nil
}
