package api

import (
	"capregsoft.com/template/models"
	api_errors "capregsoft.com/template/services/api/errors"
)

func (api *TemplateApiImpl) CreateUserId(userIdData *models.UserData) (*models.UserDataResponse, error) {
	userIdDataResp, err := api.db.CreateUserIdDB(userIdData)
	if err != nil {
		return nil, api_errors.NewInternalServerError(err.Error())
	}
	return userIdDataResp, nil
}
