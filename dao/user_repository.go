package dao

import (
	"gmarkting/models"
	"log"
	"xorm.io/xorm"
)

type SDao struct {
	engin *xorm.Engine
}

func NewSDao(engin *xorm.Engine) *SDao {

	return &SDao{
		engin: engin,
	}

}

func (d *SDao) Get(id int) *models.SysUser {
	data := &models.SysUser{Id: id}
	ok, err := d.engin.Get(data)

	if ok && err == nil {
		return data
	} else {
		return nil
	}

}

func (d *SDao) GetAll() []models.SysUser {
	//datalist := make([]models.SysUser,0)
	datalist := []models.SysUser{}

	err := d.engin.Desc("id").Find(&datalist)
	if err != nil {
		log.Fatal(err)
		return datalist
	} else {
		return datalist
	}

}

func (d *SDao) Search(country string) []models.SysUser {
	datalist := []models.SysUser{}

	err := d.engin.Where("country=?", country).Desc("id").Find(&datalist)
	if err != nil {
		log.Fatal(err)
		return datalist
	} else {
		return datalist
	}
}

func (d *SDao) Delete(id int) error {
	data := &models.SysUser{Id: id, Status: 1}

	_, err := d.engin.ID(data.Id).Update(data)

	return err

}

func (d *SDao) Update(data *models.SysUser, colums []string) error {
	//col强制更新

	_, err := d.engin.ID(data.Id).MustCols(colums...).Update(data)
	return err

}

func (d *SDao) Create(data *models.SysUser) error {
	_, err := d.engin.Insert(data)

	return err
}
