package model

import (
	"errors"
	"user/internal/service"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Nickname string
	Account  string
	Password string
	Phone    string
	Status   string
	Avatar   string
	Deposit  float32
}

const (
	PassWordCost = 12 // 密码加密难度
)

// 加密密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// 检验密码
func (user *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (user *User) CheckUserIdExist(Id string) error {
	if err := DB.Where("id=?", Id).First(&user).Error; err == gorm.ErrRecordNotFound {
		return err
	}
	return nil
}

func (user *User) CheckUserExist(Account string) bool {
	if err := DB.Where("account=?", Account).First(&user).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (user *User) ShowUserInfo(req *service.UserRequest) (err error) {
	if exist := user.CheckUserExist(req.Account); exist {
		return nil
	}
	return errors.New("Account Not Exist")
}

func BuildUser(item User) *service.UserModel {
	userModel := service.UserModel{
		Id:        uint32(item.ID),
		Account:   item.Account,
		Name:      item.Name,
		NickName:  item.Nickname,
		Phone:     item.Phone,
		CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
		Deposit:   item.Deposit,
		Status:    item.Status,
		Avatar:    item.Avatar,
	}
	return &userModel
}

func (user *User) SaveByService(userModel *service.UserModel) error {
	if exist := user.CheckUserExist(userModel.Account); exist {
		return nil
	}
	user.Account = userModel.Account
	user.Avatar = userModel.Avatar
	user.Name = userModel.Name
	user.Nickname = userModel.NickName
	user.Phone = userModel.Phone
	if err := DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
