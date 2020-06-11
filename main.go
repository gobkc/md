package main

import (
	"fmt"
	"github.com/gobkc/md/global"
	"github.com/gobkc/md/initialization"
	"log"
)

func main() {
	gin := initialization.Gin.GetString()
	binPath := initialization.Bin.GetString()
	install := initialization.Install.GetBool()
	dbConn := initialization.DbConnectParam.GetString()

	if dbConn == "" {
		//读取配置文件中的数据库链接信息
		db := global.ReadYaml().Db.Mysql
		dbConn = fmt.Sprintf("%s:%s@%s:%d/%s", db.User, db.Password, db.Server, db.Port, db.Name)
	} else {
		//命令行已经输入数据，覆盖配置文件的数据
		global.WriteYaml(dbConn)
	}

	//安装自己到/usr/bin
	if install {
		if err := global.InstallSelf(); err != nil {
			log.Fatalln(err.Error())
		}
	}

	//将当前目录下单所有文件打包成二进制
	if binPath != "" {
		global.File2Bin(binPath)
	}

	//
	if gin != "" {
		global.CreateGin(gin)
	}
}
