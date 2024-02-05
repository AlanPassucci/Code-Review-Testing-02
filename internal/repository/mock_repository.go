package repository

import (
	"app/internal"

	"github.com/stretchr/testify/mock"
)

// MockRepository is a struct that represents a mock repository
type MockRepository struct {
	mock.Mock
}

// FindAll is a method that returns a map of all vehicles
func (m *MockRepository) FindAll() (v map[int]internal.Vehicle, err error) {
	args := m.Called()
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

// FindByColorAndYear is a method that returns a map of vehicles that match the color and fabrication year
func (m *MockRepository) FindByColorAndYear(color string, fabricationYear int) (v map[int]internal.Vehicle, err error) {
	args := m.Called(color, fabricationYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

// FindByBrandAndYearRange is a method that returns a map of vehicles that match the brand and a range of fabrication years
func (m *MockRepository) FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	args := m.Called(brand, startYear, endYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

// FindByBrand is a method that returns a map of vehicles that match the brand
func (m *MockRepository) FindByBrand(brand string) (v map[int]internal.Vehicle, err error) {
	args := m.Called(brand)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

// FindByWeightRange is a method that returns a map of vehicles that match a range of weight
func (m *MockRepository) FindByWeightRange(fromWeight float64, toWeight float64) (v map[int]internal.Vehicle, err error) {
	args := m.Called(fromWeight, toWeight)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}
