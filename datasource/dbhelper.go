package datasource

import (
	"gmarkting/conf"
	"log"
	"sync"

	"xorm.io/xorm"
)

var (
	masterEngin *xorm.Engine
	slaveEngin  *xorm.Engine
	lock        sync.Mutex
)

func InstanceMaster() *xorm.Engine {

	if masterEngin != nil {
		return masterEngin
	}

	lock.Lock()
	defer lock.Unlock()

	//防止并发重复初始化
	if masterEngin != nil {
		return masterEngin
	}

	engine, err := xorm.NewEngine("mysql", conf.MasterDSName)

	if err != nil {
		log.Fatal(err)

	} else {
		masterEngin = engine
		return masterEngin
	}

	return masterEngin
}

func InstanceSlave() *xorm.Engine {
	return slaveEngin
}
