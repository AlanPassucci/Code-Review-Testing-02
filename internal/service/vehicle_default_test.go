package service

import (
	"app/internal"
	"app/internal/repository"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCase is an interface that represents a test case
type TestCase interface {
	Arrange()
	Act()
	Assert(t *testing.T)
}

// TestCaseSetup is a struct that represents the setup for a test case
type TestCaseSetup struct {
	mockRepository *repository.MockRepository
	service        *ServiceVehicleDefault
}

// Setup is a function that returns a new instance of TestCaseSetup
func Setup() *TestCaseSetup {
	mockRepository := repository.MockRepository{}
	service := NewServiceVehicleDefault(&mockRepository)

	return &TestCaseSetup{
		mockRepository: &mockRepository,
		service:        service,
	}
}

// FindByColorAndYearTestCase is a struct that represents a test case for the FindByColorAndYear method
type FindByColorAndYearTestCase struct {
	setup            *TestCaseSetup
	name             string
	color            string
	fabricationYear  int
	expectedVehicles map[int]internal.Vehicle
	expectedError    error
	repositoryError  error
	obtainedVehicles map[int]internal.Vehicle
	obtainedError    error
	mockOnCalled     bool
	isError          bool
}

// Arrange is a method that sets up the test case
func (tc *FindByColorAndYearTestCase) Arrange() {
	tc.setup = Setup()
	tc.setup.mockRepository.On("FindByColorAndYear", tc.color, tc.fabricationYear).Return(tc.expectedVehicles, tc.repositoryError)
}

// Act is a method that executes the test case
func (tc *FindByColorAndYearTestCase) Act() {
	tc.obtainedVehicles, tc.obtainedError = tc.setup.service.FindByColorAndYear(tc.color, tc.fabricationYear)
}

// Assert is a method that asserts the test case
func (tc *FindByColorAndYearTestCase) Assert(t *testing.T) {
	if tc.isError {
		assert.Error(t, tc.obtainedError)
		assert.EqualError(t, tc.obtainedError, tc.expectedError.Error())
		assert.Nil(t, tc.obtainedVehicles)
	} else {
		assert.NoError(t, tc.obtainedError)
		assert.Equal(t, tc.expectedVehicles, tc.obtainedVehicles)
	}
	if tc.mockOnCalled {
		assert.True(t, tc.setup.mockRepository.AssertCalled(t, "FindByColorAndYear", tc.color, tc.fabricationYear))
	} else {
		assert.True(t, tc.setup.mockRepository.AssertNotCalled(t, "FindByColorAndYear", tc.color, tc.fabricationYear))
	}
}

func TestService_FindByColorAndYear(t *testing.T) {
	// Create the test cases
	testCases := []FindByColorAndYearTestCase{
		{
			// This test evaluates that FindByColorAndYear returns a map of vehicles that match the color and fabrication year
			name:            "should return a map of vehicles that match the color and fabrication year",
			color:           "Red",
			fabricationYear: 2010,
			expectedVehicles: map[int]internal.Vehicle{
				1: {
					Id: 1,
					VehicleAttributes: internal.VehicleAttributes{
						Brand:           "Ford",
						Model:           "Fiesta",
						Registration:    "ABC-1234",
						Color:           "Red",
						FabricationYear: 2010,
						Capacity:        5,
						MaxSpeed:        180,
						FuelType:        "Gasoline",
						Transmission:    "Manual",
						Weight:          1000,
						Dimensions: internal.Dimensions{
							Height: 1.5,
							Length: 4,
							Width:  1.8,
						},
					},
				},
			},
			mockOnCalled: true,
		},
	}

	// Run the test cases
	for _, testCase := range testCases {
		testCase.Arrange()
		testCase.Act()
		testCase.Assert(t)
	}
}

// FindByBrandAndYearRangeTestCase is a struct that represents a test case for the FindByBrandAndYearRange method
type FindByBrandAndYearRangeTestCase struct {
	setup            *TestCaseSetup
	name             string
	brand            string
	startYear        int
	endYear          int
	expectedVehicles map[int]internal.Vehicle
	expectedError    error
	repositoryError  error
	obtainedVehicles map[int]internal.Vehicle
	obtainedError    error
	mockOnCalled     bool
	isError          bool
}

// Arrange is a method that sets up the test case
func (tc *FindByBrandAndYearRangeTestCase) Arrange() {
	tc.setup = Setup()
	tc.setup.mockRepository.On("FindByBrandAndYearRange", tc.brand, tc.startYear, tc.endYear).Return(tc.expectedVehicles, tc.repositoryError)
}

// Act is a method that executes the test case
func (tc *FindByBrandAndYearRangeTestCase) Act() {
	tc.obtainedVehicles, tc.obtainedError = tc.setup.service.FindByBrandAndYearRange(tc.brand, tc.startYear, tc.endYear)
}

// Assert is a method that asserts the test case
func (tc *FindByBrandAndYearRangeTestCase) Assert(t *testing.T) {
	if tc.isError {
		assert.Error(t, tc.obtainedError)
		assert.EqualError(t, tc.obtainedError, tc.expectedError.Error())
		assert.Nil(t, tc.obtainedVehicles)
	} else {
		assert.NoError(t, tc.obtainedError)
		assert.Equal(t, tc.expectedVehicles, tc.obtainedVehicles)
	}
	if tc.mockOnCalled {
		assert.True(t, tc.setup.mockRepository.AssertCalled(t, "FindByBrandAndYearRange", tc.brand, tc.startYear, tc.endYear))
	} else {
		assert.True(t, tc.setup.mockRepository.AssertNotCalled(t, "FindByBrandAndYearRange", tc.brand, tc.startYear, tc.endYear))
	}
}

// TestService_FindByBrandAndYearRange is a function that tests the FindByBrandAndYearRange method
func TestService_FindByBrandAndYearRange(t *testing.T) {
	// Create the test cases
	testCases := []FindByBrandAndYearRangeTestCase{
		{
			// This test evaluates that FindByBrandAndYearRange returns a map of vehicles that match the brand and year range
			name:      "should return a map of vehicles that match the brand and year range",
			brand:     "Ford",
			startYear: 2010,
			endYear:   2012,
			expectedVehicles: map[int]internal.Vehicle{
				1: {
					Id: 1,
					VehicleAttributes: internal.VehicleAttributes{
						Brand:           "Ford",
						Model:           "Fiesta",
						Registration:    "ABC-1234",
						Color:           "Red",
						FabricationYear: 2010,
						Capacity:        5,
						MaxSpeed:        180,
						FuelType:        "Gasoline",
						Transmission:    "Manual",
						Weight:          1000,
						Dimensions: internal.Dimensions{
							Height: 1.5,
							Length: 4,
							Width:  1.8,
						},
					},
				},
				2: {
					Id: 2,
					VehicleAttributes: internal.VehicleAttributes{
						Brand:           "Ford",
						Model:           "Focus",
						Registration:    "ABC-1235",
						Color:           "Blue",
						FabricationYear: 2010,
						Capacity:        5,
						MaxSpeed:        180,
						FuelType:        "Gasoline",
						Transmission:    "Manual",
						Weight:          1100,
						Dimensions: internal.Dimensions{
							Height: 1.5,
							Length: 4,
							Width:  1.8,
						},
					},
				},
			},
			mockOnCalled: true,
		},
	}

	// Run the test cases
	for _, testCase := range testCases {
		testCase.Arrange()
		testCase.Act()
		testCase.Assert(t)
	}
}

// AverageMaxSpeedByBrandTestCase is a struct that represents a test case for the AverageMaxSpeedByBrand method
type AverageMaxSpeedByBrandTestCase struct {
	setup            *TestCaseSetup
	name             string
	brand            string
	returnedVehicles map[int]internal.Vehicle
	expectedAverage  float64
	expectedError    error
	repositoryError  error
	obtainedAverage  float64
	obtainedError    error
	mockOnCalled     bool
	isError          bool
}

// Arrange is a method that sets up the test case
func (tc *AverageMaxSpeedByBrandTestCase) Arrange() {
	tc.setup = Setup()
	tc.setup.mockRepository.On("FindByBrand", tc.brand).Return(tc.returnedVehicles, tc.repositoryError)
}

// Act is a method that executes the test case
func (tc *AverageMaxSpeedByBrandTestCase) Act() {
	tc.obtainedAverage, tc.obtainedError = tc.setup.service.AverageMaxSpeedByBrand(tc.brand)
}

// Assert is a method that asserts the test case
func (tc *AverageMaxSpeedByBrandTestCase) Assert(t *testing.T) {
	if tc.isError {
		assert.Error(t, tc.obtainedError)
		assert.EqualError(t, tc.obtainedError, tc.expectedError.Error())
		assert.Equal(t, tc.expectedAverage, tc.obtainedAverage)
	} else {
		assert.NoError(t, tc.obtainedError)
		assert.Equal(t, tc.expectedAverage, tc.obtainedAverage)
	}
	if tc.mockOnCalled {
		assert.True(t, tc.setup.mockRepository.AssertCalled(t, "FindByBrand", tc.brand))
	} else {
		assert.True(t, tc.setup.mockRepository.AssertNotCalled(t, "FindByBrand", tc.brand))
	}
}

// TestService_AverageMaxSpeedByBrand is a function that tests the AverageMaxSpeedByBrand method
func TestService_AverageMaxSpeedByBrand(t *testing.T) {
	// Create the test cases
	testCases := []AverageMaxSpeedByBrandTestCase{
		{
			// This test evaluates that AverageMaxSpeedByBrand returns the average of the max speeds of the vehicles that match the brand
			name:  "should return the average of the max speeds of the vehicles that match the brand",
			brand: "Ford",
			returnedVehicles: map[int]internal.Vehicle{
				1: {
					Id: 1,
					VehicleAttributes: internal.VehicleAttributes{
						Brand:           "Ford",
						Model:           "Fiesta",
						Registration:    "ABC-1234",
						Color:           "Red",
						FabricationYear: 2010,
						Capacity:        5,
						MaxSpeed:        180,
						FuelType:        "Gasoline",
						Transmission:    "Manual",
						Weight:          1000,
						Dimensions: internal.Dimensions{
							Height: 1.5,
							Length: 4,
							Width:  1.8,
						},
					},
				},
				2: {
					Id: 2,
					VehicleAttributes: internal.VehicleAttributes{
						Brand:           "Ford",
						Model:           "Focus",
						Registration:    "ABC-1235",
						Color:           "Blue",
						FabricationYear: 2010,
						Capacity:        5,
						MaxSpeed:        180,
						FuelType:        "Gasoline",
						Transmission:    "Manual",
						Weight:          1100,
						Dimensions: internal.Dimensions{
							Height: 1.5,
							Length: 4,
							Width:  1.8,
						},
					},
				},
			},
			expectedAverage: 180,
			mockOnCalled:    true,
		},
		{
			// This test evaluates that AverageMaxSpeedByBrand returns an error when the repository returns an error
			name:            "should return an error when the repository returns an error",
			brand:           "Ford",
			expectedError:   errors.New("error"),
			repositoryError: errors.New("error"),
			mockOnCalled:    true,
			isError:         true,
		},
		{
			// This test evaluates that AverageMaxSpeedByBrand returns 0 and ErrServiceNoVehicles when there are no vehicles that match the brand
			name:          "should return 0 and ErrServiceNoVehicles when there are no vehicles that match the brand",
			brand:         "Ford",
			expectedError: internal.ErrServiceNoVehicles,
			mockOnCalled:  true,
			isError:       true,
		},
	}

	// Run the test cases
	for _, testCase := range testCases {
		testCase.Arrange()
		testCase.Act()
		testCase.Assert(t)
	}
}

// AverageCapacityByBrandTestCase is a struct that represents a test case for the AverageCapacityByBrand method
type AverageCapacityByBrandTestCase struct {
	setup            *TestCaseSetup
	name             string
	brand            string
	returnedVehicles map[int]internal.Vehicle
	expectedCapacity int
	expectedError    error
	repositoryError  error
	obtainedCapacity int
	obtainedError    error
	mockOnCalled     bool
	isError          bool
}

// Arrange is a method that sets up the test case
func (tc *AverageCapacityByBrandTestCase) Arrange() {
	tc.setup = Setup()
	tc.setup.mockRepository.On("FindByBrand", tc.brand).Return(tc.returnedVehicles, tc.repositoryError)
}

// Act is a method that executes the test case
func (tc *AverageCapacityByBrandTestCase) Act() {
	tc.obtainedCapacity, tc.obtainedError = tc.setup.service.AverageCapacityByBrand(tc.brand)
}

// Assert is a method that asserts the test case
func (tc *AverageCapacityByBrandTestCase) Assert(t *testing.T) {
	if tc.isError {
		assert.Error(t, tc.obtainedError)
		assert.EqualError(t, tc.obtainedError, tc.expectedError.Error())
		assert.Equal(t, tc.expectedCapacity, tc.obtainedCapacity)
	} else {
		assert.NoError(t, tc.obtainedError)
		assert.Equal(t, tc.expectedCapacity, tc.obtainedCapacity)
	}
	if tc.mockOnCalled {
		assert.True(t, tc.setup.mockRepository.AssertCalled(t, "FindByBrand", tc.brand))
	} else {
		assert.True(t, tc.setup.mockRepository.AssertNotCalled(t, "FindByBrand", tc.brand))
	}
}

// TestService_AverageCapacityByBrand is a function that tests the AverageCapacityByBrand method
func TestService_AverageCapacityByBrand(t *testing.T) {
	// Create the test cases
	testCases := []AverageCapacityByBrandTestCase{
		{
			// This test evaluates that AverageCapacityByBrand returns the average capacity of the vehicles that match the brand
			name:  "should return the average capacity of the vehicles that match the brand",
			brand: "Ford",
			returnedVehicles: map[int]internal.Vehicle{
				1: {
					Id: 1,
					VehicleAttributes: internal.VehicleAttributes{
						Brand:           "Ford",
						Model:           "Fiesta",
						Registration:    "ABC-1234",
						Color:           "Red",
						FabricationYear: 2010,
						Capacity:        5,
						MaxSpeed:        180,
						FuelType:        "Gasoline",
						Transmission:    "Manual",
						Weight:          1000,
						Dimensions: internal.Dimensions{
							Height: 1.5,
							Length: 4,
							Width:  1.8,
						},
					},
				},
				2: {
					Id: 2,
					VehicleAttributes: internal.VehicleAttributes{
						Brand:           "Ford",
						Model:           "Focus",
						Registration:    "ABC-1235",
						Color:           "Blue",
						FabricationYear: 2010,
						Capacity:        5,
						MaxSpeed:        180,
						FuelType:        "Gasoline",
						Transmission:    "Manual",
						Weight:          1100,
						Dimensions: internal.Dimensions{
							Height: 1.5,
							Length: 4,
							Width:  1.8,
						},
					},
				},
			},
			expectedCapacity: 5,
			mockOnCalled:     true,
		},
		{
			// This test evaluates that AverageCapacityByBrand returns an error when the repository returns an error
			name:            "should return an error when the repository returns an error",
			brand:           "Ford",
			expectedError:   errors.New("error"),
			repositoryError: errors.New("error"),
			mockOnCalled:    true,
			isError:         true,
		},
		{
			// This test evaluates that AverageCapacityByBrand returns 0 and ErrServiceNoVehicles when there are no vehicles that match the brand
			name:          "should return 0 and ErrServiceNoVehicles when there are no vehicles that match the brand",
			brand:         "Ford",
			expectedError: internal.ErrServiceNoVehicles,
			mockOnCalled:  true,
			isError:       true,
		},
	}

	// Run the test cases
	for _, testCase := range testCases {
		testCase.Arrange()
		testCase.Act()
		testCase.Assert(t)
	}
}

// SearchByWeightRangeTestCase is a struct that represents a test case for the SearchByWeightRange method
type SearchByWeightRangeTestCase struct {
	setup                                 *TestCaseSetup
	name                                  string
	query                                 internal.SearchQuery
	ok                                    bool
	returnedVehiclesFromFindAll           map[int]internal.Vehicle
	returnedFindAllError                  error
	returnedVehiclesFromFindByWeightRange map[int]internal.Vehicle
	returnedFindByWeightRangeError        error
	expectedVehicles                      map[int]internal.Vehicle
	expectedError                         error
	mockOnCalledFindAll                   bool
	mockOnCalledFindByWeightRange         bool
	isError                               bool
}

// Arrange is a method that sets up the test case
func (tc *SearchByWeightRangeTestCase) Arrange() {
	tc.setup = Setup()
	tc.setup.mockRepository.On("FindAll").Return(tc.returnedVehiclesFromFindAll, tc.returnedFindAllError)
	tc.setup.mockRepository.On("FindByWeightRange", tc.query.FromWeight, tc.query.ToWeight).Return(tc.returnedVehiclesFromFindByWeightRange, tc.returnedFindByWeightRangeError)
}

// Act is a method that executes the test case
func (tc *SearchByWeightRangeTestCase) Act() {
	tc.expectedVehicles, tc.expectedError = tc.setup.service.SearchByWeightRange(tc.query, tc.ok)
}

// Assert is a method that asserts the test case
func (tc *SearchByWeightRangeTestCase) Assert(t *testing.T) {
	if tc.isError {
		assert.Error(t, tc.expectedError)
		assert.EqualError(t, tc.expectedError, tc.expectedError.Error())
		assert.Equal(t, tc.expectedVehicles, tc.expectedVehicles)
	} else {
		assert.NoError(t, tc.expectedError)
		assert.Equal(t, tc.expectedVehicles, tc.expectedVehicles)
	}
	if tc.mockOnCalledFindAll {
		assert.True(t, tc.setup.mockRepository.AssertCalled(t, "FindAll"))
	} else {
		assert.True(t, tc.setup.mockRepository.AssertNotCalled(t, "FindAll"))
	}
	if tc.mockOnCalledFindByWeightRange {
		assert.True(t, tc.setup.mockRepository.AssertCalled(t, "FindByWeightRange", tc.query.FromWeight, tc.query.ToWeight))
	} else {
		assert.True(t, tc.setup.mockRepository.AssertNotCalled(t, "FindByWeightRange", tc.query.FromWeight, tc.query.ToWeight))
	}
}

// TestService_SearchByWeightRange is a function that tests the SearchByWeightRange method
func TestService_SearchByWeightRange(t *testing.T) {
	// Create the test cases
	testCases := []SearchByWeightRangeTestCase{
		{
			// This test evaluates that SearchByWeightRange returns the vehicles that match the weight range
			name: "should return the vehicles that match the weight range",
			query: internal.SearchQuery{
				FromWeight: 1100,
				ToWeight:   1150,
			},
			ok: true,
			returnedVehiclesFromFindByWeightRange: map[int]internal.Vehicle{
				1: {
					Id: 1,
					VehicleAttributes: internal.VehicleAttributes{
						Brand:           "Ford",
						Model:           "Fiesta",
						Registration:    "ABC-1234",
						Color:           "Red",
						FabricationYear: 2010,
						Capacity:        5,
						MaxSpeed:        180,
						FuelType:        "Gasoline",
						Transmission:    "Manual",
						Weight:          1000,
						Dimensions: internal.Dimensions{
							Height: 1.5,
							Length: 4,
							Width:  1.8,
						},
					},
				},
			},
			expectedVehicles: map[int]internal.Vehicle{
				1: {
					Id: 1,
					VehicleAttributes: internal.VehicleAttributes{
						Brand:           "Ford",
						Model:           "Fiesta",
						Registration:    "ABC-1234",
						Color:           "Red",
						FabricationYear: 2010,
						Capacity:        5,
						MaxSpeed:        180,
						FuelType:        "Gasoline",
						Transmission:    "Manual",
						Weight:          1000,
						Dimensions: internal.Dimensions{
							Height: 1.5,
							Length: 4,
							Width:  1.8,
						},
					},
				},
			},
			mockOnCalledFindByWeightRange: true,
		},
		{
			// This test evaluates that SearchByWeightRange returns all the vehicles when the ok parameter is false
			name: "should return all the vehicles when the ok parameter is false",
			returnedVehiclesFromFindAll: map[int]internal.Vehicle{
				1: {
					Id: 1,
					VehicleAttributes: internal.VehicleAttributes{
						Brand:           "Ford",
						Model:           "Fiesta",
						Registration:    "ABC-1234",
						Color:           "Red",
						FabricationYear: 2010,
						Capacity:        5,
						MaxSpeed:        180,
						FuelType:        "Gasoline",
						Transmission:    "Manual",
						Weight:          1000,
						Dimensions: internal.Dimensions{
							Height: 1.5,
							Length: 4,
							Width:  1.8,
						},
					},
				},
				2: {
					Id: 2,
					VehicleAttributes: internal.VehicleAttributes{
						Brand:           "Ford",
						Model:           "Focus",
						Registration:    "ABC-1235",
						Color:           "Blue",
						FabricationYear: 2010,
						Capacity:        5,
						MaxSpeed:        180,
						FuelType:        "Gasoline",
						Transmission:    "Manual",
						Weight:          1100,
						Dimensions: internal.Dimensions{
							Height: 1.5,
							Length: 4,
							Width:  1.8,
						},
					},
				},
				3: {
					Id: 3,
					VehicleAttributes: internal.VehicleAttributes{
						Brand:           "Fiat",
						Model:           "Uno",
						Registration:    "ABC-1236",
						Color:           "Red",
						FabricationYear: 2012,
						Capacity:        5,
						MaxSpeed:        180,
						FuelType:        "Gasoline",
						Transmission:    "Manual",
						Weight:          1200,
						Dimensions: internal.Dimensions{
							Height: 1.5,
							Length: 4,
							Width:  1.8,
						},
					},
				},
			},
			expectedVehicles: map[int]internal.Vehicle{
				1: {
					Id: 1,
					VehicleAttributes: internal.VehicleAttributes{
						Brand:           "Ford",
						Model:           "Fiesta",
						Registration:    "ABC-1234",
						Color:           "Red",
						FabricationYear: 2010,
						Capacity:        5,
						MaxSpeed:        180,
						FuelType:        "Gasoline",
						Transmission:    "Manual",
						Weight:          1000,
						Dimensions: internal.Dimensions{
							Height: 1.5,
							Length: 4,
							Width:  1.8,
						},
					},
				},
				2: {
					Id: 2,
					VehicleAttributes: internal.VehicleAttributes{
						Brand:           "Ford",
						Model:           "Focus",
						Registration:    "ABC-1235",
						Color:           "Blue",
						FabricationYear: 2010,
						Capacity:        5,
						MaxSpeed:        180,
						FuelType:        "Gasoline",
						Transmission:    "Manual",
						Weight:          1100,
						Dimensions: internal.Dimensions{
							Height: 1.5,
							Length: 4,
							Width:  1.8,
						},
					},
				},
				3: {
					Id: 3,
					VehicleAttributes: internal.VehicleAttributes{
						Brand:           "Fiat",
						Model:           "Uno",
						Registration:    "ABC-1236",
						Color:           "Red",
						FabricationYear: 2012,
						Capacity:        5,
						MaxSpeed:        180,
						FuelType:        "Gasoline",
						Transmission:    "Manual",
						Weight:          1200,
						Dimensions: internal.Dimensions{
							Height: 1.5,
							Length: 4,
							Width:  1.8,
						},
					},
				},
			},
			mockOnCalledFindAll: true,
		},
	}

	// Run the test cases
	for _, testCase := range testCases {
		testCase.Arrange()
		testCase.Act()
		testCase.Assert(t)
	}
}
