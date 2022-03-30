package db

import (
	"fmt"

	"capregsoft.com/template/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type TemplateDB interface {
	CreateUserIdDB(userData *models.UserData) (*models.UserDataResponse, error)
	GetUserDataByIdDB(userId int) (*models.UserDataResponse, error)
}
type TemplateDBImpl struct {
	dbConn *sqlx.DB
}

func NewTemplateDBImpl() *TemplateDBImpl {
	dbConn, err := sqlx.Connect("mysql", "database:Aswad_database@123@tcp(127.0.0.1:3306)/template")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("database is connected")
	}
	return &TemplateDBImpl{
		dbConn: dbConn,
	}
}

var _ TemplateDB = &TemplateDBImpl{}
