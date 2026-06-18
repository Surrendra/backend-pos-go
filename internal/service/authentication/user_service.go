package service

import (
	"BackendPOS/internal/constant"
	"BackendPOS/internal/helper"
	model "BackendPOS/internal/model/authentication"
	authRepository "BackendPOS/internal/repository/authentication"
	maintenanceRepository "BackendPOS/internal/repository/maintenance"
	baseRequest "BackendPOS/internal/request"
	authRequest "BackendPOS/internal/request/authentication"
	"BackendPOS/internal/response"
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

func NewUserService(
	userRepository authRepository.UserRepositoryInterface,
	roleRepository authRepository.RoleRepositoryInterface,
	merchantRepository maintenanceRepository.MerchantRepositoryInterface,
) *UserService {
	return &UserService{
		userRepository:     userRepository,
		roleRepository:     roleRepository,
		merchantRepository: merchantRepository,
	}
}

func (s *UserService) Login(username string, password string) (*model.User, error) {
	user, err := s.userRepository.FindByUsernameOrEmail(username)
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

func (s *UserService) Register(req authRequest.RegisterRequest) (*model.User, error) {
	var user model.User
	user.Name = req.Name
	user.Code = uuid.NewString()
	user.Email = req.Email
	user.Username = req.Username
	user.Address = &req.Address
	user.Phone = &req.Phone
	user.Active = constant.IndicatorActive
	user.Status = constant.StatusActive

	defaultRole, err := s.roleRepository.FindByName(constant.RoleOperatorName)
	if err != nil {
		return nil, fmt.Errorf("default role not found, please contact administrator")
	}
	defaultMerchant, err := s.merchantRepository.FindByCode(constant.MerchantSampleCode)
	if err != nil {
		return nil, fmt.Errorf("default merchant not found, please contact administrator")
	}
	user.RoleId = defaultRole.ID
	user.MerchantId = defaultMerchant.ID
	user.MerchantCode = &defaultMerchant.Code
	user.MerchantName = &defaultMerchant.Name
	user.Password, _ = helper.HashPassword(req.Password)
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

func (s *UserService) Create(request authRequest.CreateUserRequest) (*model.User, error) {
	var user model.User
	user.Name = request.Name
	user.Code = uuid.NewString()
	user.Email = request.Email
	user.Username = request.Username
	user.Address = &request.Address
	user.Phone = &request.Phone
	user.Active = constant.IndicatorActive
	user.Status = constant.StatusActive

	defaultRole, err := s.roleRepository.FindByName(request.RoleCodes[0])
	if err != nil {
		return nil, fmt.Errorf("default role not found, please contact administrator")
	}
	defaultMerchant, err := s.merchantRepository.FindByCode(constant.MerchantSampleCode)
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

func (s *UserService) GetData() ([]model.User, error) {
	users, err := s.userRepository.GetData()
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}
	return users, nil
}

func (s *UserService) FindByCode(code string) (*model.User, error) {
	user, err := s.userRepository.FindByCode(code)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by code: %v", err)
	}
	return &user, nil
}

func (s *UserService) Update(code string, req authRequest.UpdateUserRequest) (*model.User, error) {
	findUser, err := s.userRepository.FindByCode(code)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by code: %v", err)
	}
	var user model.User
	user.Name = req.Name
	user.Email = req.Email
	user.Username = req.Username
	user.Active = req.Active
	user.Status = req.Status
	if req.Address != "" {
		user.Address = &req.Address
	}
	if req.Phone != "" {
		user.Phone = &req.Phone
	}
	if req.Password != "" {
		user.Password, _ = helper.HashPassword(req.Password)
	}
	updatedUser, err := s.userRepository.Update(findUser.Code, user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}
	if len(req.RoleCodes) > 0 {
		newRole, err := s.roleRepository.FindByName(req.RoleCodes[0])
		if err == nil && newRole.ID != 0 {
			updateRole := model.User{RoleId: newRole.ID}
			_, _ = s.userRepository.Update(findUser.Code, updateRole)
			updatedUser.RoleId = newRole.ID
		}
	}
	return &updatedUser, nil
}

func (s *UserService) Delete(code string) error {
	findUser, err := s.userRepository.FindByCode(code)
	if err != nil {
		return fmt.Errorf("failed to find user by code: %v", err)
	}
	err = s.userRepository.Delete(findUser.Code)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}

func (s *UserService) Datatable(req baseRequest.DataTableRequest) ([]model.User, response.DataTableMeta, error) {
	return s.userRepository.Datatable(req)
}
