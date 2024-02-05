package repository

import (
	"app/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCase is an interface that represents a test case
type TestCase interface {
	Act()
	Assert(t *testing.T)
}

// TestCaseSetup is a struct that represents a test case setup
type TestCaseSetup struct {
	repository *RepositoryReadVehicleMap
}

// Setup is a function that returns a new instance of TestCaseSetup
func Setup() *TestCaseSetup {
	db := map[int]internal.Vehicle{
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
	}
	return &TestCaseSetup{
		repository: NewRepositoryReadVehicleMap(db),
	}
}

// FindAllTestCase is a struct that represents a test case for the FindAll method
type FindAllTestCase struct {
	setup            *TestCaseSetup
	name             string
	expectedVehicles map[int]internal.Vehicle
	expectedError    error
	obtainVehicles   map[int]internal.Vehicle
	obteinedError    error
	isError          bool
}

// Act is a method that executes the test case
func (tc *FindAllTestCase) Act() {
	tc.obtainVehicles, tc.obteinedError = tc.setup.repository.FindAll()
}

// Assert is a method that asserts the test case
func (tc *FindAllTestCase) Assert(t *testing.T) {
	if tc.isError {
		assert.Error(t, tc.obteinedError)
		assert.EqualError(t, tc.obteinedError, tc.expectedError.Error())
		assert.Nil(t, tc.obtainVehicles)
	} else {
		assert.NoError(t, tc.obteinedError)
		assert.Equal(t, tc.expectedVehicles, tc.obtainVehicles)
	}
}

// TestRepository_FindAll is a test function that tests the FindAll method
func TestRepository_FindAll(t *testing.T) {
	// Create the test cases
	testCases := []FindAllTestCase{
		{
			// This test case evaluates that FindAll returns a map of vehicles
			name: "should return a map of vehicles",
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
		},
	}

	// Run the test cases
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setup = Setup()
			testCase.Act()
			testCase.Assert(t)
		})
	}
}

// FindByColorAndYearTestCase is a struct that represents a test case for the FindByColorAndYear method
type FindByColorAndYearTestCase struct {
	setup            *TestCaseSetup
	name             string
	color            string
	year             int
	expectedVehicles map[int]internal.Vehicle
	expectedError    error
	obtainVehicles   map[int]internal.Vehicle
	obteinedError    error
	isError          bool
}

// Act is a method that executes the test case
func (tc *FindByColorAndYearTestCase) Act() {
	tc.obtainVehicles, tc.obteinedError = tc.setup.repository.FindByColorAndYear(tc.color, tc.year)
}

// Assert is a method that asserts the test case
func (tc *FindByColorAndYearTestCase) Assert(t *testing.T) {
	if tc.isError {
		assert.Error(t, tc.obteinedError)
		assert.EqualError(t, tc.obteinedError, tc.expectedError.Error())
		assert.Nil(t, tc.obtainVehicles)
	} else {
		assert.NoError(t, tc.obteinedError)
		assert.Equal(t, tc.expectedVehicles, tc.obtainVehicles)
	}
}

// TestRepository_FindByColorAndYear is a test function that tests the FindByColorAndYear method
func TestRepository_FindByColorAndYear(t *testing.T) {
	// Create the test cases
	testCases := []FindByColorAndYearTestCase{
		{
			// This test case evaluates that FindByColorAndYear returns a map of vehicles that match the color and year
			name:  "should return a map of vehicles that match the color and year",
			color: "Red",
			year:  2010,
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
		},
	}

	// Run the test cases
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setup = Setup()
			testCase.Act()
			testCase.Assert(t)
		})
	}
}

// FindByBrandAndYearRangeTestCase is a struct that represents a test case for the FindByBrandAndYearRange method
type FindByBrandAndYearRangeTestCase struct {
	setup            *TestCaseSetup
	name             string
	brand            string
	starYear         int
	endYear          int
	expectedVehicles map[int]internal.Vehicle
	expectedError    error
	obtainVehicles   map[int]internal.Vehicle
	obteinedError    error
	isError          bool
}

// Act is a method that executes the test case
func (tc *FindByBrandAndYearRangeTestCase) Act() {
	tc.obtainVehicles, tc.obteinedError = tc.setup.repository.FindByBrandAndYearRange(tc.brand, tc.starYear, tc.endYear)
}

// Assert is a method that asserts the test case
func (tc *FindByBrandAndYearRangeTestCase) Assert(t *testing.T) {
	if tc.isError {
		assert.Error(t, tc.obteinedError)
		assert.EqualError(t, tc.obteinedError, tc.expectedError.Error())
		assert.Nil(t, tc.obtainVehicles)
	} else {
		assert.NoError(t, tc.obteinedError)
		assert.Equal(t, tc.expectedVehicles, tc.obtainVehicles)
	}
}

// TestRepository_FindByBrandAndYearRange is a test function that tests the FindByBrandAndYearRange method
func TestRepository_FindByBrandAndYearRange(t *testing.T) {
	// Create the test cases
	testCases := []FindByBrandAndYearRangeTestCase{
		{
			// This test case evaluates that FindByBrandAndYearRange returns a map of vehicles that match the brand and year range
			name:     "should return a map of vehicles that match the brand and year range",
			brand:    "Ford",
			starYear: 2010,
			endYear:  2012,
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
		},
	}

	// Run the test cases
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setup = Setup()
			testCase.Act()
			testCase.Assert(t)
		})
	}
}

// FindByBrandTestCase is a struct that represents a test case for the FindByBrand method
type FindByBrandTestCase struct {
	setup            *TestCaseSetup
	name             string
	brand            string
	expectedVehicles map[int]internal.Vehicle
	expectedError    error
	obtainVehicles   map[int]internal.Vehicle
	obteinedError    error
	isError          bool
}

// Act is a method that executes the test case
func (tc *FindByBrandTestCase) Act() {
	tc.obtainVehicles, tc.obteinedError = tc.setup.repository.FindByBrand(tc.brand)
}

// Assert is a method that asserts the test case
func (tc *FindByBrandTestCase) Assert(t *testing.T) {
	if tc.isError {
		assert.Error(t, tc.obteinedError)
		assert.EqualError(t, tc.obteinedError, tc.expectedError.Error())
		assert.Nil(t, tc.obtainVehicles)
	} else {
		assert.NoError(t, tc.obteinedError)
		assert.Equal(t, tc.expectedVehicles, tc.obtainVehicles)
	}
}

// TestRepository_FindByBrand is a test function that tests the FindByBrand method
func TestRepository_FindByBrand(t *testing.T) {
	// Create the test cases
	testCases := []FindByBrandTestCase{
		{
			// This test case evaluates that FindByBrand returns a map of vehicles that match the brand
			name:  "should return a map of vehicles that match the brand",
			brand: "Fiat",
			expectedVehicles: map[int]internal.Vehicle{
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
		},
	}

	// Run the test cases
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setup = Setup()
			testCase.Act()
			testCase.Assert(t)
		})
	}
}

// FindByWeightRangeTestCase is a struct that represents a test case for the FindByWeightRange method
type FindByWeightRangeTestCase struct {
	setup            *TestCaseSetup
	name             string
	fromWeight       float64
	toWeight         float64
	expectedVehicles map[int]internal.Vehicle
	expectedError    error
	obtainVehicles   map[int]internal.Vehicle
	obteinedError    error
	isError          bool
}

// Act is a method that executes the test case
func (tc *FindByWeightRangeTestCase) Act() {
	tc.obtainVehicles, tc.obteinedError = tc.setup.repository.FindByWeightRange(tc.fromWeight, tc.toWeight)
}

// Assert is a method that asserts the test case
func (tc *FindByWeightRangeTestCase) Assert(t *testing.T) {
	if tc.isError {
		assert.Error(t, tc.obteinedError)
		assert.EqualError(t, tc.obteinedError, tc.expectedError.Error())
		assert.Nil(t, tc.obtainVehicles)
	} else {
		assert.NoError(t, tc.obteinedError)
		assert.Equal(t, tc.expectedVehicles, tc.obtainVehicles)
	}
}

// TestRepository_FindByWeightRange is a test function that tests the FindByWeightRange method
func TestRepository_FindByWeightRange(t *testing.T) {
	// Create the test cases
	testCases := []FindByWeightRangeTestCase{
		{
			// This test case evaluates that FindByWeightRange returns a map of vehicles that match the weight range
			name:       "should return a map of vehicles that match the weight range",
			fromWeight: 1100,
			toWeight:   1200,
			expectedVehicles: map[int]internal.Vehicle{
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
		},
	}

	// Run the test cases
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setup = Setup()
			testCase.Act()
			testCase.Assert(t)
		})
	}
}
