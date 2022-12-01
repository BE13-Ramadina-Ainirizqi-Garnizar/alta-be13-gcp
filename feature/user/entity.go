package user

import "time"

type Core struct {
	ID        uint
	Name      string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
	Phone     string `validate:"required"`
	Address   string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) (row int, err error)
	GetByID(id int) (data []Core, err error)
	Delete(data Core, id int) (row int, err error)
	Update(input Core, id int) (row int, err error)
}

type ServiceInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) (err error)
	GetByID(id int) (data []Core, err error)
	Delete(data Core, id int) (err error)
	Update(input Core, id int) (err error)
}
