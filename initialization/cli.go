package initialization

import (
	CmdParse "github.com/gobkc/cmd-parse"
)

var Install = CmdParse.New().SetItem("install").SetUsage("install(install self)").SetDefault(false).SaveSet()
var Bin = CmdParse.New().SetItem("create-bin").SetUsage("create-bin [your dir name]").SetDefault("dir_name").SaveSet()
var Gin = CmdParse.New().SetItem("create").SetUsage("create [your project name]").SetDefault("project_name").SaveSet()
var Swagger = CmdParse.New().SetItem("-swagger").SetUsage("-swagger(only -swagger)").SetDefault(false).SaveSet()
var Db = CmdParse.New().SetItem("create-db").SetUsage("create-db [your database name]").SetDefault("dbname").SaveSet()
var Model = CmdParse.New().SetItem("create-model").SetUsage("create-model [your database name]").SetDefault("dbname").SaveSet()
var FromTable = CmdParse.New().SetItem("-from").SetUsage("-from [your table name]").SetDefault("table_name").SaveSet()
var Migrate = CmdParse.New().SetItem("create-migrate").SetUsage("create-migrate [your migrate name]").SetDefault("dbname").SaveSet()
var Seed = CmdParse.New().SetItem("create-seed").SetUsage("create-migrate [your seed name]").SetDefault("seed_name").SaveSet()
var MigrateDo = CmdParse.New().SetItem("run-migrate").SetUsage("run-migrate [your migrate name]").SetDefault("test_migrate").SaveSet()
var SeedDo = CmdParse.New().SetItem("run-seed").SetUsage("run-seed [your seed name]").SetDefault("seed_name").SaveSet()
var DbConnectParam = CmdParse.New().SetItem("-s").SetUsage("s server:db name:user:password").SetDefault("").SaveSet()
var HasExplain = CmdParse.New().Explain("注意事项").
	SetExplainItem("部分功能须在网络联通环境下使用").
	SetExplainItem("数据迁移和数据填充设计用于在测试环境/开发环境使用").
	Explain("使用示例").
	SetExplainItem("md create-bin your dir").
	SetExplainItem("md create my_project").
	SetExplainItem("md create-model user -from user -s root:123456@127.0.0.1:3306/my_db").
	SetExplainItem("md create-db my_db -s root:123456@127.0.0.1:3306/").
	SetExplainItem("md create-migrate my_migrate -s root:123456@127.0.0.1:3306/my_db").
	SetExplainItem("md create-seed my_seed -s root:123456@127.0.0.1:3306/my_db").
	SetExplainItem("md run-seed my_seed -s root:123456@127.0.0.1:3306/my_db").
	SetExplainItem("md run-migrate my_migrate -s root:123456@127.0.0.1:3306/my_db").
	SaveExplain()
