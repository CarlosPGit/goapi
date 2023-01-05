package admin

import (
	"fmt"

	"github.com/CarlosPGit/golang_sface/internal/admin/user"
	jwt_config "github.com/CarlosPGit/golang_sface/internal/platform/jwtConfig"
	"golang.org/x/crypto/bcrypt"
)

const errorFindUser = "user doesnt exists"

type UserUseCase interface {
	Login(user user.User) (string, error)
	Register(user user.User) bool
}

type Admin struct {
	userRepo *user.UserRepo
}

func NewUserUseCase(ur *user.UserRepo) *Admin {
	return &Admin{userRepo: ur}
}

func (adm *Admin) Login(user user.User) (string, error) {
	var token string

	userBD := adm.userRepo.FindByEmail(user.Email)
	if userBD.ID > 0 {
		errCheck := checkPasswordHash(user, userBD.Password)
		if errCheck == nil {
			return jwt_config.GenerateToken(user)
		}

		return token, errCheck
	}

	return token, fmt.Errorf(errorFindUser)

}

func (adm *Admin) Register(user user.User) bool {
	if user.ID > 0 {
		return false
	}

	if adm.userRepo.FindByEmail(user.Email).ID > 0 {
		fmt.Println("Email exists")
		return false
	}

	hashPassword(&user)

	return adm.userRepo.Save(user)
}

func checkPasswordHash(user user.User, passwordHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(user.Password))
	return err
}

func hashPassword(user *user.User) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		panic(err)
	}

	user.Password = string(bytes)
}
