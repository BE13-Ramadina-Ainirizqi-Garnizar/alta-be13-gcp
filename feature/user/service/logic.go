package service

import (
	"Ramadina/CleanArchitecture/feature/user"
	"errors"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userRepository user.RepositoryInterface
	validate       *validator.Validate
}

func New(repo user.RepositoryInterface) user.ServiceInterface {
	return &userService{
		userRepository: repo,
		validate:       validator.New(),
	}
}

func (s *userService) GetAll() (data []user.Core, err error) {
	data, err = s.userRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *userService) Create(input user.Core) (err error) {
	// if input.Name == "" || input.Email == "" || input.Password == "" {
	// 	return errors.New("nama, email dan password harus diisi")
	// }

	if errValidate := s.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	input.Role = "user"

	_, errCreate := s.userRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

func (s *userService) GetByID(id int) (data []user.Core, err error) {
	data, err = s.userRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *userService) Delete(user user.Core, id int) error {
	_, errDelete := s.userRepository.Delete(user, id)
	if errDelete != nil {
		return errors.New("error delete")
	}

	return nil
}

func (s *userService) Update(user user.Core, id int) error {
	_, errUpdate := s.userRepository.Update(user, id)
	if errUpdate != nil {
		return errors.New("error update")
	}
	return nil
}
