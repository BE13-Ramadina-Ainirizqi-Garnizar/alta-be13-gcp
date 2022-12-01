package main

import (
	"Ramadina/CleanArchitecture/config"
	"Ramadina/CleanArchitecture/factory"
	"Ramadina/CleanArchitecture/utils/database/mysql"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)

	e := echo.New()

	factory.InitFactory(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}
