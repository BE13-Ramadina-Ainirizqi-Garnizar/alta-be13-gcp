package service

import (
	"Ramadina/CleanArchitecture/feature/user"
	"Ramadina/CleanArchitecture/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	repo := new(mocks.UserRepository)
	returnData := []user.Core{{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}}
	t.Run("Success get all data", func(t *testing.T) {
		repo.On("GetAll").Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Name, response[0].Name)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get All data", func(t *testing.T) {
		repo.On("GetAll").Return(nil, errors.New("failed to get data")).Once()
		srv := New(repo)
		response, err := srv.GetAll()
		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	repo := new(mocks.UserRepository)
	t.Run("Success Create user", func(t *testing.T) {
		inputRepo := user.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}
		inputData := user.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta"}
		repo.On("Create", inputRepo).Return(1, nil).Once()
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Create user, duplicate entry", func(t *testing.T) {
		inputRepo := user.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}
		inputData := user.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta"}
		repo.On("Create", inputRepo).Return(0, errors.New("failed to insert data, error query")).Once()
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "failed to insert data, error query", err.Error())
		repo.AssertExpectations(t)
	})

	t.Run("Failed Create user, name empty", func(t *testing.T) {
		inputData := user.Core{Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta"}
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	repo := new(mocks.UserRepository)
	returnData := []user.Core{{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}}
	t.Run("Sucess read data", func(t *testing.T) {
		inputID := 1
		repo.On("GetByID", inputID).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetByID(inputID)
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Name, response[0].Name)
	})

	t.Run("Error read data", func(t *testing.T) {
		inputID := 1
		repo.On("GetByID", inputID).Return(nil, errors.New("failed to get data")).Once()
		srv := New(repo)
		response, err := srv.GetByID(inputID)
		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mocks.UserRepository)
	t.Run("Success delete user", func(t *testing.T) {
		inputData := user.Core{}
		inputID := 1
		repo.On("Delete", inputData, inputID).Return(1, nil).Once()
		srv := New(repo)
		err := srv.Delete(inputData, inputID)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Delete User", func(t *testing.T) {
		inputData := user.Core{}
		inputID := 1
		repo.On("Delete", inputData, inputID).Return(0, errors.New("failed to delete data")).Once()
		srv := New(repo)
		err := srv.Delete(inputData, inputID)
		assert.NotNil(t, err)
		assert.Equal(t, "error delete", err.Error())
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := new(mocks.UserRepository)
	t.Run("Success update user", func(t *testing.T) {
		inputRepo := user.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}
		// inputData := user.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta"}
		inputID := 1
		repo.On("Update", inputRepo, inputID).Return(1, nil).Once()
		srv := New(repo)
		err := srv.Update(inputRepo, inputID)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Update User", func(t *testing.T) {
		inputRepo := user.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}
		inputID := 1
		repo.On("Update", inputRepo, inputID).Return(0, errors.New("failed to update data")).Once()
		srv := New(repo)
		err := srv.Update(inputRepo, inputID)
		assert.NotNil(t, err)
		assert.Equal(t, "error update", err.Error())
		repo.AssertExpectations(t)
	})
}
