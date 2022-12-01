package factory

import (
	authDelivery "Ramadina/CleanArchitecture/feature/auth/delivery"
	authRepo "Ramadina/CleanArchitecture/feature/auth/repository"
	authService "Ramadina/CleanArchitecture/feature/auth/service"

	userDelivery "Ramadina/CleanArchitecture/feature/user/delivery"
	userRepo "Ramadina/CleanArchitecture/feature/user/repository"
	userService "Ramadina/CleanArchitecture/feature/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userRepoFactory := userRepo.New(db)
	// userRepoFactory := userRepo.NewRaw(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, e)

	authRepoFactory := authRepo.New(db)
	authServiceFactory := authService.New(authRepoFactory)
	authDelivery.New(authServiceFactory, e)

}
