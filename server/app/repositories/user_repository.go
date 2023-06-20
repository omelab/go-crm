// app/repositories/user_repository.go
package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/omelab/go-crm/app/models"
	"golang.org/x/crypto/bcrypt"
)

 
type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	ComparePasswords(hashedPassword string, plainPassword string) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	if err := r.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) ComparePasswords(hashedPassword string, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}