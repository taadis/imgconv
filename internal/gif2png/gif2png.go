package gif2png

import (
	"image/gif"
	"image/png"

	"os"
)

// ConvertGIFToPNG 将指定的 GIF 文件转换为 PNG 文件
func ConvertGIFToPNG(inputPath string, outputPath string) error {
	// 打开 GIF 文件
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 解码 GIF 图像
	gifImage, err := gif.DecodeAll(file)
	if err != nil {
		return err
	}

	// 如果 GIF 包含多帧，这里只处理第一帧
	// 如果需要处理所有帧，可以选择保存为多张 PNG 或其他方式
	pngImage := gifImage.Image[0]

	// 创建输出文件
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// 编码为 PNG 格式
	err = png.Encode(outFile, pngImage)
	if err != nil {
		return err
	}

	return nil
}
