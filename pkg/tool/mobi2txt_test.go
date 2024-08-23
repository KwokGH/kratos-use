package tool

import (
	"fmt"
	"github.com/unidoc/unioffice/document"
	"os"
	"testing"
)

func TestMobi2txt(t *testing.T) {
	Mobi2Txt()
}

func Mobi2Txt() {
	p := "C:\\Users\\user\\Documents\\WeChat Files\\wxid_4901619006212\\FileStorage\\File\\2024-06\\安徒生童话.mobi"
	// 打开mobi文件
	doc, err := document.Open(p)
	if err != nil {
		fmt.Println("Error opening document:", err)
		return
	}

	// 读取文档内容并转换为txt
	text := doc.ExtractText().Text()

	// 将内容写入txt文件
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(text)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}

	fmt.Println("Conversion successful. Check output.txt.")
}
