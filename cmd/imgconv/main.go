package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/taadis/imgconv/internal/gif2png"
)

func main() {
	inputPattern := flag.String("input", "", "输入的 GIF 文件路径或模式")
	outputPattern := flag.String("output", "", "输出的 PNG 文件路径或模式")
	flag.Parse()

	fmt.Printf("输入参数: input=%s, output=%s\n", *inputPattern, *outputPattern)

	if *inputPattern == "" {
		log.Fatal("请使用 -input 参数指定 GIF 文件的路径或模式")
	}
	if *outputPattern == "" {
		log.Fatal("请使用 -output 参数指定 PNG 文件的路径或模式")
	}

	inputFiles, err := filepath.Glob(*inputPattern)
	if err != nil {
		log.Fatalf("无法匹配输入文件:%v", err)
	}
	log.Printf("获得输入文件数=%d", len(inputFiles))
	if len(inputFiles) == 0 {
		log.Fatal("没有匹配的输入文件")
	}

	outputDir := filepath.Dir(*outputPattern)
	for _, inputPath := range inputFiles {
		outputFileName := strings.Replace(filepath.Base(inputPath), ".gif", ".png", 1)
		outputPath := filepath.Join(outputDir, outputFileName)
		err := gif2png.ConvertGIFToPNG(inputPath, outputPath)
		if err != nil {
			log.Fatalf("转换失败: %v", err)
		}
		fmt.Printf("成功将 %s 转换为 %s\n", inputPath, outputPath)
	}
}
