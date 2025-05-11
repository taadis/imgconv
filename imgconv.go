package imgconv

import "io"

// Converter 定义了图片格式转换的接口
type Converter interface {
	// Convert 将输入图片转换为目标格式
	// input 是输入图片数据的读取器
	// output 是输出图片数据的写入器
	// 返回转换过程中的错误，如果转换成功则返回nil
	Convert(input io.Reader, output io.Writer) error
}
