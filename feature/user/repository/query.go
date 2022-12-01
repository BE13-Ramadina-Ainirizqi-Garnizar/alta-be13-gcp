package repository

import (
	"Ramadina/CleanArchitecture/feature/user"
	"errors"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.RepositoryInterface {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetAll() (data []user.Core, err error) {
	var users []User

	tx := r.db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	dataCore := toCoreList(users)
	return dataCore, nil
}

func (r *userRepository) Create(input user.Core) (row int, err error) {
	userGorm := fromCore(input)
	tx := r.db.Create(&userGorm)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil
}

func (r *userRepository) GetByID(id int) (data []user.Core, err error) {
	var user []User

	tx := r.db.First(&user, id)

	dataCore := toCoreList(user)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return dataCore, nil
}

func (r *userRepository) Delete(data user.Core, id int) (row int, err error) {
	userGorm := fromCore(data)

	tx := r.db.Delete(&userGorm, id)
	if tx.Error != nil {
		return -1, err
	}

	if tx.RowsAffected == 0 {
		return 0, err
	}

	return int(tx.RowsAffected), nil
}

func (r *userRepository) Update(data user.Core, id int) (row int, err error) {
	var user User
	gormUserCore := fromCore(data)

	tx := r.db.First(&user, id)

	if tx.Error != nil {
		return -1, err
	}

	tz := r.db.Model(&user).Updates(gormUserCore)
	if tz.Error != nil {
		return -1, err
	}

	if tz.RowsAffected == 0 {
		return 0, errors.New("error insert")
	}

	return int(tz.RowsAffected), nil
}
