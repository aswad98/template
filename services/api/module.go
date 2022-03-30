package api

import (
	"capregsoft.com/template/models"
	"capregsoft.com/template/services/db"
)

type TemplateApi interface {
	CreateUserId(userIdData *models.UserData) (*models.UserDataResponse, error)
	GetUserDataById(userId int) (*models.UserDataResponse, error)
}
type TemplateApiImpl struct {
	db *db.TemplateDBImpl
}

func NewTemplateApiImpl() *TemplateApiImpl {
	dbImpl := db.NewTemplateDBImpl()
	return &TemplateApiImpl{
		db: dbImpl,
	}
}

var _ TemplateApi = &TemplateApiImpl{}
