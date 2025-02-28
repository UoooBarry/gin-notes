package services

import (
	"uooobarry/gin-notes/models"
	"uooobarry/gin-notes/repositories"

	"golang.org/x/crypto/bcrypt"
)

// UserService 业务逻辑层
type UserService struct {
	UserRepo *repositories.UserRepository
}

// NewUserService 创建一个新的 UserService
func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

// CreateUser 创建一个新的用户并保存
func (service *UserService) CreateUser(username, password string) (*models.User, error) {
	// 密码加密
	user := &models.User{
		Username: username,
	}
	err := user.SetPassword(password)
	if err != nil {
		return nil, err
	}

	// 保存用户到数据库
	err = service.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *UserService) AuthenticateUser(username, password string) (*models.User, error) {
	user, err := service.UserRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	// 校验密码
	if !user.CheckPassword(password) {
		return nil, bcrypt.ErrMismatchedHashAndPassword
	}

	return user, nil
}
