package repository

import (
	_user "Ramadina/CleanArchitecture/feature/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Phone    string `gorm:"type:varchar(15)"`
	Address  string
	Role     string
}

func fromCore(dataCore _user.Core) User {
	userGorm := User{
		Name:     dataCore.Name,
		Email:    dataCore.Email,
		Password: dataCore.Password,
		Phone:    dataCore.Phone,
		Address:  dataCore.Address,
		Role:     dataCore.Role,
	}
	return userGorm
}

func (dataModel *User) toCore() _user.Core {
	return _user.Core{
		ID:        dataModel.ID,
		Name:      dataModel.Name,
		Email:     dataModel.Email,
		Password:  dataModel.Password,
		Phone:     dataModel.Phone,
		Address:   dataModel.Address,
		Role:      dataModel.Role,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

func toCoreList(dataModel []User) []_user.Core {
	var dataCore []_user.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
