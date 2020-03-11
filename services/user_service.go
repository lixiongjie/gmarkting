package services

import (
	"gmarkting/dao"
	"gmarkting/datasource"
	"gmarkting/models"
)

type SService interface {
	GetAll() []models.SysUser
	Get(id int) *models.SysUser
	Delete(id int) error

	Update(user *models.SysUser, colums []string) error
	Create(user *models.SysUser) error

	Search(country string) []models.SysUser
}

type sService struct {
	dao *dao.SDao
}

func NewSService() SService {

	return &sService{
		dao: dao.NewSDao(datasource.InstanceMaster()),
	}

}

func (s *sService) GetAll() []models.SysUser {
	return s.dao.GetAll()
}

func (s *sService) Get(id int) *models.SysUser {
	return s.dao.Get(id)
}

func (s *sService) Delete(id int) error {
	return s.dao.Delete(id)
}

func (s *sService) Update(user *models.SysUser, colums []string) error {
	return s.dao.Update(user, colums)
}

func (s *sService) Create(user *models.SysUser) error {
	return s.dao.Create(user)
}

func (s *sService) Search(country string) []models.SysUser {
	return s.dao.Search(country)
}
