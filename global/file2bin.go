package global

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const Exclude = "vendor;.git;.idea;"
const BinInputDir = "."

func File2Bin(binOutDir string) {
	dir := BinInputDir
	if dir[len(dir)-1:] == "/" {
		dir = dir[len(dir)-1:]
	}
	var dirLen = len(dir)
	var fList []string
	var err error

	if err = MakeInitFile(binOutDir); err != nil {
		log.Println(err.Error())
	}

	if fList, err = GetAllFile(dir, fList); err != nil {
		log.Fatalln(err.Error())
	}
	outFile := time.Now().Format("2006-01-02 15:04:05 File")
	for i, filePath := range fList {
		bytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			continue
		}
		filePath = filePath[dirLen:]
		outFileName := fmt.Sprintf("%s/%s%d.go", binOutDir, outFile, i)
		byteData := Byte2String(bytes)
		outData := fmt.Sprintf("package %s\n\nfunc init() {\n\tFileTree[\"%s\"] =\"%s\"\n}", binOutDir, filePath, string(byteData))
		if _, err := os.Create(outFileName); err != nil {
			log.Println(err.Error())
			continue
		}
		if err := ioutil.WriteFile(outFileName, []byte(outData), 0644); err != nil {
			os.RemoveAll(outFileName)
		}
	}
}

func MakeInitFile(binOutDir string) error {
	//判断输出目录是否存在
	if _, err := os.Stat(binOutDir); err != nil && os.IsNotExist(err) {
		if err := os.Mkdir(binOutDir, os.ModePerm); err != nil {
			log.Fatalln("创建输出目录失败，详细：", err.Error())
		}
	}
	outFileName := fmt.Sprintf("%s/init.go", binOutDir)
	if _, err := os.Create(outFileName); err != nil {
		return err
	}
	outData := fmt.Sprintf("package %s\n\nvar FileTree = make(map[string]string)", binOutDir)
	if err := ioutil.WriteFile(outFileName, []byte(outData), 0644); err != nil {
		os.RemoveAll(outFileName)
	}
	return nil
}

func Byte2String(b []byte) string {
	var result string
	var tmp []string
	for _, v := range b {
		tmp = append(tmp, fmt.Sprintf("%v", v))
	}
	result = strings.Join(tmp, " ")
	return result
}

func String2Byte(s string) []byte {
	var ori []string
	var tmp []byte
	ori = strings.Split(s, " ")
	for _, v := range ori {
		if newv, err := strconv.Atoi(v); err == nil {
			tmp = append(tmp, uint8(newv))
		}
	}
	return tmp
}

func GetAllFile(pathname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}
	for _, fi := range rd {
		if isExclude := strings.Contains(Exclude, fi.Name()); isExclude {
			continue
		}
		if fi.IsDir() {
			fullDir := pathname + "/" + fi.Name()
			s, err = GetAllFile(fullDir, s)
			if err != nil {
				fmt.Println("read dir fail:", err)
				return s, err
			}
		} else {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}
