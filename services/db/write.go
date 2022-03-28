package db

import (
	db_errors "capregsoft.com/template/services/db/errors"

	"capregsoft.com/template/models"
)

func (db *TemplateDBImpl) CreateUserIdDB(userData *models.UserData) (*models.UserDataResponse, error) {
	userIdResp := models.NewUserDataResponse()
	tx := db.dbConn.MustBegin()
	_, err := tx.NamedQuery(`INSERT INTO userdata(user_id,name,email,phone_number)
	VALUES(:user_id,:name,:email,:phone_number)`, userData)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, db_errors.NewInternalServerError(err.Error())
	}
	err = db.dbConn.Get(userIdResp, `SELECT * FROM userdata WHERE user_id=?;`, *userData.UserId)
	if err != nil {
		return nil, db_errors.NewInternalServerError(err.Error())
	}
	return userIdResp, nil
}
