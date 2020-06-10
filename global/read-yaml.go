package global

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
)

type MdConf struct {
	Db struct {
		Mysql struct {
			Name     string `yaml:"name"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			Server   string `yaml:"server"`
			Port     int    `yaml:"port"`
		} `yaml:"mysql"`
	}
}

const ConfFile = "md-conf.yaml"

func ReadYaml() (result MdConf) {
	var data []byte
	var err error
	if data, err = ioutil.ReadFile(ConfFile); err != nil {
		log.Println(err.Error())
		return result
	}
	mdConf := MdConf{}
	err = yaml.Unmarshal([]byte(data), &mdConf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return mdConf
}

func WriteYaml(dbConn string) (result string) {
	if !PathExists(ConfFile) {
		if fn, err := os.Create(ConfFile); err != nil {
			log.Println(err.Error())
		} else {
			fn.Close()
		}
	}
	var mdConf = ReadYaml()
	var cond = regexp.MustCompile(`(.*):(.*)@(.*):(.*)/(.*)`)
	var findMatch = cond.FindStringSubmatch(dbConn)
	var dbUser, dbPass, dbServer, dbPort, dbName = findMatch[1], findMatch[2], findMatch[3], findMatch[4], findMatch[5]
	var portInt, _ = strconv.Atoi(dbPort)
	mdConf.Db.Mysql = struct {
		Name     string `yaml:"name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Server   string `yaml:"server"`
		Port     int    `yaml:"port"`
	}{Name: dbName, User: dbUser, Password: dbPass, Server: dbServer, Port: portInt}
	if yarByte, err := yaml.Marshal(&mdConf); err != nil {
		log.Println(err.Error())
	} else {
		ioutil.WriteFile(ConfFile, yarByte, 0644)
	}
	return result
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	log.Println(err.Error())
	return false
}
