package db

import (
	"capregsoft.com/template/models"
	db_errors "capregsoft.com/template/services/db/errors"
)

func (db *TemplateDBImpl) GetUserDataByIdDB(userId int) (*models.UserDataResponse, error) {
	userDataDB := models.NewUserDataResponse()
	err := db.dbConn.Get(userDataDB, `SELECT * FROM userdata WHERE user_id=?d`, userId)
	if err != nil {
		return nil, db_errors.NewError("No data found against this id")
	}
	return userDataDB, nil

}
