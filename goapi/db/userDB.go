package db

import "gorm.io/gorm"

// type UserDB struct {
// 	gorm.Model
// 	ID    uint `gorm:"primary_key"`
// 	Name  string
// 	Email string
// 	Todo  string
// }

type UserDB struct {
	gorm.Model
	ID    uint `gorm:"primary_key"`
	Name  string
	Email string
	Todo  string
}

func All() (*[]UserDB, error) {
	db := DB()

	var userDB []UserDB
	if err := db.Find(&userDB).Error; err != nil {
		return nil, err
	}

	return &userDB, nil
}

func Bymail(email string) (*[]UserDB, error) {
	db := DB()

	var userDB []UserDB
	if err := db.Where(&UserDB{Email: email}).Find(&userDB).Error; err != nil {
		return nil, err
	}

	return &userDB, nil
}

func Read(id string) (*UserDB, error) {
	db := DB()

	var userDB UserDB
	if err := db.First(&userDB, id).Error; err != nil {
		return nil, err
	}

	return &userDB, nil
}

func Create(user UserDB) error {
	db := DB()
	userData := UserDB{
		Name:  user.Name,
		Email: user.Email,
		Todo:  user.Todo,
	}

	if err := db.Create(&userData).Error; err != nil {
		return err
	}
	return nil
}

// func Read(id string) (*UserDB, error)
func Update(user UserDB) error {
	db := DB()
	userData := UserDB{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Todo:  user.Todo,
	}

	if err := db.Updates(&userData).Error; err != nil {
		return err
	}
	return nil
}

func Delete(id string) error {
	db := DB()

	var userDB UserDB
	if err := db.Delete(&userDB, id).Error; err != nil {
		return err
	}

	return nil
}
