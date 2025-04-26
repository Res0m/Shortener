package user

import "GolangAdvanced/pkg/db"

// CREATE
// FindByEmail

type UserRepository struct {
	DataBase *db.Db
}


func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		DataBase: database,
	}
}


func (repo *UserRepository) Create(user *User) (*User, error){
	result := repo.DataBase.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*User, error){
	var userF User
	result := repo.DataBase.DB.First(&userF, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userF, nil
}

