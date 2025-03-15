package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/taadis/imgconv/internal/gif2png"
)

func main() {
	inputPath := flag.String("input", "", "输入的 GIF 文件路径")
	outputPath := flag.String("output", "output.png", "输出的 PNG 文件路径")
	flag.Parse()

	if *inputPath == "" {
		log.Fatal("请使用 -input 参数指定 GIF 文件的路径")
	}

	if _, err := os.Stat(*inputPath); os.IsNotExist(err) {
		log.Fatalf("指定的文件 %s 不存在", *inputPath)
	}

	if err := gif2png.ConvertGIFToPNG(*inputPath, *outputPath); err != nil {
		log.Fatalf("转换失败: %v", err)
	}

	fmt.Printf("成功将 %s 转换为 %s\n", *inputPath, *outputPath)
}
