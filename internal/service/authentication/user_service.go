package service

import (
	constan "BackendPOS/internal/constant"
	"BackendPOS/internal/helper"
	model "BackendPOS/internal/model/authentication"
	authRepository "BackendPOS/internal/repository/authentication"
	maintenanceRepository "BackendPOS/internal/repository/maintenance"
	request "BackendPOS/internal/request/authentication"
	"BackendPOS/internal/util"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository     authRepository.UserRepositoryInterface
	roleRepository     authRepository.RoleRepositoryInterface
	merchantRepository maintenanceRepository.MerchantRepositoryInterface
}

func NewUserService(userRepository authRepository.UserRepositoryInterface, roleRepository authRepository.RoleRepositoryInterface, merchantRepository maintenanceRepository.MerchantRepositoryInterface) *UserService {
	return &UserService{userRepository: userRepository, roleRepository: roleRepository, merchantRepository: merchantRepository}
}

func (s *UserService) Login(username string, password string) (*model.User, error) {
	user, err := s.userRepository.FindByUsername(username)
	if err != nil {
		//return nil, fmt.Errorf("Something wrong when find user by username %s", username)
		return nil, fmt.Errorf("User with username %s does not exist", username)
	}
	// if user not found
	if user.ID == 0 {
		return nil, fmt.Errorf("User with username %s does not exist", username)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("Username and password do not match")
	}
	token, err := util.GenerateToken(user.Code, user.ID, user.Username, user.Name, user.MerchantId, user.RoleId)
	if err != nil {
		return nil, fmt.Errorf("Something wrong when generate token")
	}

	// add token to user
	user.Token = &token
	return user, nil
}

func (s *UserService) Register(request request.RegisterRequest) (*model.User, error) {
	var user model.User
	user.Name = request.Name
	user.Code = uuid.NewString()
	user.Email = request.Email
	user.Username = request.Username
	user.Address = &request.Address
	user.Phone = &request.Phone
	user.Active = constan.IndicatorActive
	user.Status = constan.StatusActive

	defaultRole, err := s.roleRepository.FindByName(constan.RoleOperatorName)
	if err != nil {
		return nil, fmt.Errorf("default role not found, please contact administrator")
	}
	defaultMerchant, err := s.merchantRepository.FindByCode(constan.MerchantSampleCode)
	if err != nil {
		return nil, fmt.Errorf("default merchant not found, please contact administrator")
	}
	user.RoleId = defaultRole.ID
	user.MerchantId = defaultMerchant.ID
	user.MerchantCode = &defaultMerchant.Code
	user.MerchantName = &defaultMerchant.Name
	user.Password, _ = helper.HashPassword(request.Password)
	newUser, err := s.userRepository.Create(user)
	if err != nil {
		return nil, fmt.Errorf("Something wrong when create user")
	}
	s.merchantRepository.CreateDirect(newUser.ID, defaultMerchant.ID, newUser.ID, newUser.Name)
	s.roleRepository.CreateDirect(newUser.ID, defaultRole.ID, newUser.ID, newUser.Name)
	token, err := util.GenerateToken(newUser.Code, newUser.ID, newUser.Username, newUser.Name, newUser.MerchantId, newUser.RoleId)
	newUser.Token = &token
	return &newUser, nil

}
