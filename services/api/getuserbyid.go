package api

import (
	"capregsoft.com/template/models"
	api_errors "capregsoft.com/template/services/api/errors"
)

func (api *TemplateApiImpl) GetUserDataById(userId int) (*models.UserDataResponse, error) {
	getUserData, err := api.db.GetUserDataByIdDB(userId)
	if err != nil {
		return nil, api_errors.NewInternalServerError(err.Error())
	}
	return getUserData, nil
}
