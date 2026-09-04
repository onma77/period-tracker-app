package repository

import "cycle-app-backend/internal/models"

type UserRepository struct {
	users  []models.User
	nextID int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  []models.User{},
		nextID: 1,
	}
}

func (r *UserRepository) Create(user models.User) models.User {
	user.ID = r.nextID
	r.nextID++

	r.users = append(r.users, user)

	return user
}

func (r *UserRepository) GetByEmail(email string) (*models.User, bool) {
	for _, user := range r.users {
		if user.Email == email {
			return &user, true
		}
	}
	return nil, false
}

func (r *UserRepository) GetAll() []models.User {
	return r.users
}
