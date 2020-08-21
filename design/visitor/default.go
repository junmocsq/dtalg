package visitor

import (
	"fmt"
	"strings"
)

type resourceFile interface {
	accept(v visitor)
	getFilePath() string
}

type pptFile struct {
	filePath string
}

func NewPPTFile(f string) resourceFile {
	return &pptFile{filePath: f}
}

func (f *pptFile) accept(v visitor) {
	v.visit(f)
}
func (f *pptFile) getFilePath() string {
	return f.filePath
}

type wordFile struct {
	filePath string
}

func NewWordFile(f string) resourceFile {
	return &wordFile{filePath: f}
}
func (f *wordFile) accept(v visitor) {
	v.visit(f)
}
func (f *wordFile) getFilePath() string {
	return f.filePath
}

type execlFile struct {
	filePath string
}

func NewExeclFile(f string) resourceFile {
	return &execlFile{filePath: f}
}
func (f *execlFile) accept(v visitor) {
	v.visit(f)
}

func (f *execlFile) getFilePath() string {
	return f.filePath
}

func toolApplication() {
	fileArr := []string{
		"a.ppt",
		"b.word",
		"c.execl",
		"d.ppt",
		"e.execl",
	}
	// 提取数据
	list := listAllResourceFile(fileArr)
	for _, resource := range list {
		resource.accept(NewExtractor())
		resource.accept(NewCompressor())
		resource.accept(NewA())
		fmt.Println()
	}

	// 压缩数据
	//for _, resource := range list {
	//	resource.accept(NewCompressor())
	//}
	//
	//for _, resource := range list {
	//	resource.accept(NewA())
	//}
}

// 初始化数据列表
func listAllResourceFile(fileArr []string) []resourceFile {
	list := make([]resourceFile, 0)
	for _, v := range fileArr {
		if strings.HasSuffix(v, ".ppt") {
			list = append(list, NewPPTFile(v))
		} else if strings.HasSuffix(v, ".word") {
			list = append(list, NewWordFile(v))
		} else if strings.HasSuffix(v, ".execl") {
			list = append(list, NewExeclFile(v))
		}

	}
	return list
}
