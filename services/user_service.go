package services

import "iris/models"

type SService interface {
	GetAll() []models.SysUser
	Get(id int64) (models.SysUser, bool)
	Delete(id int64) bool

	Update(user *models.SysUser) error
	Create(user *models.SysUser) error

	Search(country string) []models.SysUser
}
