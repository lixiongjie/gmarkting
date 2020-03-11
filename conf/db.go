package conf

const MasterDSName = "root:123456@tcp(127.0.0.1:3306)/?charset=utf8"

type DbConf struct {
	Host   string
	Port   int
	User   string
	Pwd    string
	DbName string
}

var MasterDbConfig DbConf = DbConf{
	Host:   "127.0.0.1",
	Port:   3306,
	User:   "root",
	Pwd:    "root",
	DbName: "cus",
}

var SlavebConfig DbConf = DbConf{
	Host:   "127.0.0.1",
	Port:   3307,
	User:   "root",
	Pwd:    "root",
	DbName: "cus",
}
