package mock

import (
	"photon-server/pkg/model"

	"github.com/stretchr/testify/mock"
)

// UserRepository is mock
type UserRepository struct {
	mock.Mock
}

// GetAll is mock function
func (_m *UserRepository) GetAll() ([]model.UserTable, error) {
	ret := _m.Called()
	return ret.Get(0).([]model.UserTable), ret.Error(1)
}

// GetByEmailAndPassword is mock function
func (_m *UserRepository) GetByEmailAndPassword(email string, password string) (model.UserTable, error) {
	ret := _m.Called(email, password)
	return ret.Get(0).(model.UserTable), ret.Error(1)
}

// Create is mock function
func (_m *UserRepository) Create(user model.User) (int64, error) {
	ret := _m.Called(user)
	return ret.Get(0).(int64), ret.Error(1)
}

// Update is mock function
func (_m *UserRepository) Update(user model.User) error {
	ret := _m.Called(user)
	return ret.Error(0)
}

// Delete is mock function
func (_m *UserRepository) Delete(id int) error {
	ret := _m.Called(id)
	return ret.Error(0)
}
