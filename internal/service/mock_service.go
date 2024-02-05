package service

import (
	"app/internal"

	"github.com/stretchr/testify/mock"
)

// MockService is a struct that implements the Service interface
type MockService struct {
	mock.Mock
}

// FindByColorAndYear is a method that returns a map of vehicles that match the color and fabrication year
func (m *MockService) FindByColorAndYear(color string, fabricationYear int) (v map[int]internal.Vehicle, err error) {
	args := m.Called(color, fabricationYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

// FindByBrandAndYearRange is a method that returns a map of vehicles that match the brand and a range of fabrication years
func (m *MockService) FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	args := m.Called(brand, startYear, endYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

// AverageMaxSpeedByBrand is a method that returns the average speed of the vehicles by brand
func (m *MockService) AverageMaxSpeedByBrand(brand string) (a float64, err error) {
	args := m.Called(brand)
	return args.Get(0).(float64), args.Error(1)
}

// AverageCapacityByBrand is a method that returns the average capacity of the vehicles by brand
func (m *MockService) AverageCapacityByBrand(brand string) (a int, err error) {
	args := m.Called(brand)
	return args.Get(0).(int), args.Error(1)
}

// SearchByWeightRange is a method that returns a map of vehicles that match the weight range
func (m *MockService) SearchByWeightRange(query internal.SearchQuery, ok bool) (v map[int]internal.Vehicle, err error) {
	args := m.Called(query, ok)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}
