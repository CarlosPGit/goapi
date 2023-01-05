package handlers

import (
	"fmt"

	"github.com/CarlosPGit/golang_sface/internal/admin"
	"github.com/CarlosPGit/golang_sface/internal/admin/user"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	useCase admin.UserUseCase
}

func NewUserHandler(useCase admin.UserUseCase) *User {
	return &User{
		useCase: useCase,
	}
}

func (u User) Login(c *fiber.Ctx) error {
	var user user.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	token, err := u.useCase.Login(user)
	if err != nil {
		fmt.Print(err.Error())
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.JSON(fiber.Map{"token": token})
}

func (u User) Register(c *fiber.Ctx) error {
	var user user.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(u.useCase.Register(user))
}
