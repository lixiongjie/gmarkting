package datasource

import (
	"sync"

	"xorm.io/xorm"
)

var (
	masterEngin *xorm.Engine
	slaveEngin  *xorm.Engine
	lock        sync.Mutex
)

func InstanceMaster() *xorm.Engine {
	return masterEngin
}

func InstanceSlave() *xorm.Engine {
	return slaveEngin
}
