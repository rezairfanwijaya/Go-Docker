package user

import "gorm.io/gorm"

// interface
type IUserRepository interface {
	FindByEmail(email string) (User, error)
	FindAll() ([]User, error)
	Save(user User) (User, error)
}

type UserRepository struct {
	db *gorm.DB
}

// new repo
func NewRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

// implementasi
func (r *UserRepository) Save(user User) (User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (User, error) {
	var user User
	if err := r.db.Where("email = ? ", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) FindAll() ([]User, error) {
	var user []User
	if err := r.db.Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
