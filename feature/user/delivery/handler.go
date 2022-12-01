package delivery

import (
	"Ramadina/CleanArchitecture/feature/user"
	"Ramadina/CleanArchitecture/middlewares"
	"Ramadina/CleanArchitecture/utils/helper"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userService user.ServiceInterface
}

func New(service user.ServiceInterface, e *echo.Echo) {
	handler := &UserDelivery{
		userService: service,
	}

	e.GET("/users", handler.GetAll, middlewares.JWTMiddleware())
	e.POST("/users", handler.Create)
	e.GET("/users/:id", handler.GetByID)
	e.DELETE("/users/:id", handler.Delete)
	e.PUT("/users/:id", handler.Update)
}

func (d *UserDelivery) GetAll(c echo.Context) error {
	results, err := d.userService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all data users", dataResponse))
}

func (d *UserDelivery) Create(c echo.Context) error {
	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := toCore(userInput)
	err := d.userService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success create data"))
}

func (d *UserDelivery) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	result, err := d.userService.GetByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data", result))
}

func (d *UserDelivery) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := user.Core{}

	errResult := d.userService.Delete(user, id)
	if errResult != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error delete data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))
}

func (d *UserDelivery) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := user.Core{}

	c.Bind(&user)

	resultErr := d.userService.Update(user, id)

	if resultErr != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error update data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update data"))
}
