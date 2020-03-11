package models

import (
	"time"
)

type SysUser struct {
	Id          int       `xorm:"not null pk autoincr comment('用户id') INT(11)"`
	Username    string    `xorm:"not null default '' comment('用户名称') VARCHAR(20)"`
	Telephone   string    `xorm:"not null default '' comment('手机号') VARCHAR(13)"`
	Mail        string    `xorm:"not null default '' comment('邮箱') VARCHAR(20)"`
	Password    string    `xorm:"not null default '' comment('加密后的密码') VARCHAR(40)"`
	DeptId      int       `xorm:"not null default 0 comment('用户所在部门的id') INT(11)"`
	Status      int       `xorm:"not null default 1 comment('状态，1：正常，0：冻结状态，2：删除') INT(11)"`
	Remark      string    `xorm:"default '' comment('备注') VARCHAR(200)"`
	Operator    string    `xorm:"not null default '' comment('操作者') VARCHAR(20)"`
	OperateTime time.Time `xorm:"not null comment('最后一次更新时间') DATETIME"`
	OperateIp   string    `xorm:"not null default '' comment('最后一次更新者的ip地址') VARCHAR(20)"`
}
