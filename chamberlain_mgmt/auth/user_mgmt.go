package auth

import (
	"chamberlain_mgmt/config"
	"chamberlain_mgmt/log"
	"errors"
)

type User struct {
	Username    string `form:"Username" json:"Username" binding:"required" gorm:"column:USERNAME"`
	Password    string `gorm:"column:PASSWORD"`
	Role        string `gorm:"column:ROLE"`
	NewPassword string `gorm:"-"`
}

type UserMgmt interface {
	AddUser() error
	UpdateUser() error
	DeleteUser() error
	GetUser() error
	GetUsersCount() (int64, error)
	GetUsers(offset int, limit int) ([]User, error)
	CheckAuth() (bool, error)
	ResetPassword() error
}

func (User) TableName() string  {
	return "USERS"
}

func (user *User) Adduser() error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.Create(&user)
	return result.Error
}

func (user *User) UpdateUser() error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.Model(&User{}).Where("USERNAME = ?", user.Username).Update("ROLE", user.Role).Update("PASSWORD", user.Password)

	return result.Error
}

func (user *User) DeleteUser() error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.Delete(&user, "USERNAME = ?", user.Username)
	return result.Error
}

func (user *User) GetUser() error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.Select("USERNAME", "ROLE").Find(&user, "USERNAME = ?", user.Username)
	return result.Error
}

func (user *User) GetUsersCount() (int64, error) {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return 0, errors.New("database connection is nil")
	}
	var count int64
	result := db.Model(&user).Count(&count)
	return count, result.Error
}

func (user *User) GetUsers(offset int, limit int) ([]User, error) {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return nil, errors.New("database connection is nil")
	}
	users := make([]User, 0)
	result := db.Select("USERNAME", "ROLE").Limit(limit).Offset(offset).Find(&users)
	return users, result.Error
}

func (user *User) CheckAuth() (bool, error) {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return false, errors.New("database connection is nil")
	}
	result := db.Select("USERNAME", "ROLE").Find(&user, "USERNAME = ? AND PASSWORD = ?", user.Username, user.Password)
	log.Info("username %s's role is %s", user.Username, user.Role)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, errors.New("username or password is not right")
	}
	return true, nil
}

func (user *User) ResetPassword() error {
	isPassRight, err := user.CheckAuth()
	if !isPassRight || err != nil {
		log.Warn("error password, not change password")
		return err
	}
	user.Password = user.NewPassword
	err = user.UpdateUser()
	return err
}
