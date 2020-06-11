package global

import (
	"fmt"
	"github.com/gobkc/md/gin"
	"io/ioutil"
	"os"
	"path"
)

type GinReplace struct {
	PackName string
}

//创建Gin
func CreateGin(packName string) error {
	ginTree := gin.FileTree
	var ginReplace = GinReplace{
		PackName: packName,
	}
	for mKey, rawData := range ginTree {
		content := String2Byte(rawData)
		if newContent, err := ParseHtml(string(content), ginReplace); err != nil {
			return err
		} else {
			if err := OutPutProject(packName, newContent, mKey); err != nil {
				return err
			}
		}
	}

	return nil
}

//输出创建项目文件
func OutPutProject(packName string, data string, filePath string) error {
	//1.递归创建文件夹
	basePath := fmt.Sprintf(".%s%s%s", string(os.PathSeparator), packName, path.Dir(filePath))
	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		return err
	}

	//2.准备待会儿要写入的数据
	var byteData = []byte(data)

	//3.写入文件
	filePath = fmt.Sprintf(".%s%s%s", string(os.PathSeparator), packName, filePath)
	if err := ioutil.WriteFile(filePath, byteData, 0644); err != nil {
		return err
	}
	return nil
}
