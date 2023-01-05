package api

import (
	"github.com/CarlosPGit/golang_sface/cmd/app/api/handlers"
	"github.com/CarlosPGit/golang_sface/internal/admin"
	"github.com/CarlosPGit/golang_sface/internal/admin/user"
	mysqlclient "github.com/CarlosPGit/golang_sface/internal/platform/mysql_client"
	"github.com/gofiber/fiber/v2"
)

type InitConf struct {
	userHandler *handlers.User
}

func StartApp() {
	app := fiber.New()

	handlers := start()

	MapUrls(app, handlers)

	app.Listen(":3000")
}

func start() *InitConf {
	db, err := mysqlclient.Connect()
	if err != nil {
		panic(err)
	}

	userRepo := user.NewUserRepo(db)
	userUseCase := admin.NewUserUseCase(userRepo)

	return &InitConf{
		userHandler: handlers.NewUserHandler(userUseCase),
	}
}
