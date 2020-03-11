package repositories

import (
	"errors"
	"sync"

	"xorm.io/xorm"
)

type SDao struct {
	engin *xorm.Engine
}
