package auth

import (
	"chamberlain_mgmt/config"
	"chamberlain_mgmt/log"
	"errors"
)

type User struct {
	Username string `form:"Username" json:"Username" binding:"required"`
	Password string
	Role     string
}

type UserMgmt interface {
	AddUser() error
	UpdateUser() error
	DeleteUser() error
	GetUser() error
	GetUsersCount() (int64, error)
	GetUsers(offset int, limit int) ([]User, error)
	CheckAuth() (bool, error)
}

func (user *User) Adduser() error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("There is no user named " + user.Username)
	}
	return nil

}

func (user *User) UpdateUser() error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.Model(&User{}).Where("username = ?", user.Username).Update("Password", user.Password)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("There is no user named " + user.Username)
	}
	return nil

}

func (user *User) DeleteUser() error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.Delete(&user, "Username = ?", user.Username)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("There is no user named " + user.Username)
	}
	return nil
}

func (user *User) GetUser() error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.Select("Username", "Role").Find(&user, "Username = ?", user.Username)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("There is no user named " + user.Username)
	}
	return nil
}

func (user *User) GetUsersCount() (int64, error) {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return 0, errors.New("database connection is nil")
	}
	var count int64
	result := db.Model(&user).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("There is no user named " + user.Username)
	}
	return count, nil
}

func (user *User) GetUsers(offset int, limit int) ([]User, error) {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return nil, errors.New("database connection is nil")
	}
	users := make([]User, 0)
	result := db.Select("Username", "Role").Limit(limit).Offset(offset).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("There is no user named " + user.Username)
	}
	return users, nil
}

func (user *User) CheckAuth() (bool, error) {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return false, errors.New("database connection is nil")
	}
	result := db.Select("Username", "Role").Find(&user, "Username = ? and Password = ?", user.Username, user.Password)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, errors.New("username or password is not right")
	}
	log.Info("username %s's role is %s", user.Username, user.Role)
	return true, nil
}