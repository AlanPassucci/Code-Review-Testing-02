package handler

import (
	"app/internal"
	"app/internal/service"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// basePath is the base path for the handlers
var basePath = "/vehicles"

// TestCase is an interface that represents a test case
type TestCase interface {
	Arrange()
	Assert(t *testing.T)
}

// TestCaseServerSetup is a struct that represents a server setup for a test case
type TestCaseServerSetup struct {
	server      *gin.Engine
	mockContext mock.AnythingOfTypeArgument
	mockService *service.MockService
	handler     *HandlerVehicle
}

// Setup is a function that configures and returns a TestCaseServerSetup
func Setup() *TestCaseServerSetup {
	server := gin.New()
	mockContext := mock.AnythingOfType("*gin.Context")
	mockService := service.MockService{}
	handler := NewHandlerVehicle(&mockService)

	return &TestCaseServerSetup{
		server:      server,
		mockContext: mockContext,
		mockService: &mockService,
		handler:     handler,
	}
}

// TestCaseHttpSetup is a struct that represents a http setup for a test case
type TestCaseHttpSetup struct {
	isErrorResponse    bool
	expectedStatusCode int
	expectedHeaders    http.Header
	expectedResponse   []byte
	req                *http.Request
	res                *httptest.ResponseRecorder
}

// FindByColorAndYearTestCase is a struct that represents a test case for FindByColorAndYear
type FindByColorAndYearTestCase struct {
	setup            *TestCaseServerSetup
	name             string
	color            string
	year             interface{}
	successMessage   string
	returnedVehicles map[int]internal.Vehicle
	serviceError     error
	handlerError     error
	mockOnCalled     bool
	httpSetup        *TestCaseHttpSetup
}

// Arrange is a method that sets up the test case
func (tc *FindByColorAndYearTestCase) Arrange() {
	tc.setup = Setup()
	tc.setup.server.GET(basePath+"/color/:color/year/:year", tc.setup.handler.FindByColorAndYear())

	tc.httpSetup.expectedHeaders = http.Header{
		"Content-Type": []string{"application/json"},
	}
	var expectedResponse interface{}
	if !tc.httpSetup.isErrorResponse {
		expectedResponse = map[string]interface{}{
			"message": tc.successMessage,
			"data":    tc.returnedVehicles,
		}
	} else {
		expectedResponse = map[string]interface{}{
			"status":  http.StatusText(tc.httpSetup.expectedStatusCode),
			"message": tc.handlerError.Error(),
		}
	}
	tc.httpSetup.expectedResponse, _ = json.Marshal(expectedResponse)

	tc.setup.mockService.On("FindByColorAndYear", tc.color, tc.year).Return(tc.returnedVehicles, tc.serviceError)

	tc.httpSetup.req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/color/%s/year/%v", basePath, tc.color, tc.year), nil)
	tc.httpSetup.res = httptest.NewRecorder()
}

// Assert is a method that asserts the test case
func (tc *FindByColorAndYearTestCase) Assert(t *testing.T) {
	if tc.mockOnCalled {
		assert.True(t, tc.setup.mockService.AssertCalled(t, "FindByColorAndYear", tc.color, tc.year))
	} else {
		assert.True(t, tc.setup.mockService.AssertNotCalled(t, "FindByColorAndYear", tc.color, tc.year))
	}
	assert.Equal(t, tc.httpSetup.expectedStatusCode, tc.httpSetup.res.Code)
	assert.Equal(t, tc.httpSetup.expectedHeaders, tc.httpSetup.res.Header())
	assert.JSONEq(t, string(tc.httpSetup.expectedResponse), tc.httpSetup.res.Body.String())
}

func TestHandler_FindByColorAndYear(t *testing.T) {
	// Create test cases
	testCases := []FindByColorAndYearTestCase{
		{
			// This test evaluates that FindByColorAndYear returns 200 ok and a map of vehicles that match the color and fabrication year
			name:           "should return 200 ok and a map of vehicles that match the color and fabrication year",
			color:          "Red",
			year:           2010,
			successMessage: "vehicles found",
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
			},
			mockOnCalled: true,
			httpSetup: &TestCaseHttpSetup{
				expectedStatusCode: http.StatusOK,
			},
		},
		{
			// This test evaluates that FindByColorAndYear returns a 400 bad request error when the year is invalid
			name:         "should return a 400 bad request error when the year is invalid",
			color:        "Red",
			year:         "a",
			handlerError: errors.New("invalid year"),
			httpSetup: &TestCaseHttpSetup{
				isErrorResponse:    true,
				expectedStatusCode: http.StatusBadRequest,
			},
		},
		{
			// This test evaluates that FindByColorAndYear returns 500 internal server error when the service returns an error
			name:         "should return 500 internal server error when the service returns an error",
			color:        "Red",
			year:         2010,
			serviceError: errors.New("error"),
			handlerError: errors.New("internal error"),
			mockOnCalled: true,
			httpSetup: &TestCaseHttpSetup{
				isErrorResponse:    true,
				expectedStatusCode: http.StatusInternalServerError,
			},
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		testCase.Arrange()
		testCase.setup.server.ServeHTTP(testCase.httpSetup.res, testCase.httpSetup.req)
		testCase.Assert(t)
	}
}

// FindByBrandAndYearRangeTestCase is a struct that represents a test case for FindByBrandAndYearRange
type FindByBrandAndYearRangeTestCase struct {
	setup            *TestCaseServerSetup
	name             string
	brand            string
	startYear        interface{}
	endYear          interface{}
	successMessage   string
	returnedVehicles map[int]internal.Vehicle
	serviceError     error
	handlerError     error
	mockOnCalled     bool
	httpSetup        *TestCaseHttpSetup
}

// Arrange is a method that sets up the test case
func (tc *FindByBrandAndYearRangeTestCase) Arrange() {
	tc.setup = Setup()
	tc.setup.server.GET(basePath+"/brand/:brand/between/:start_year/:end_year", tc.setup.handler.FindByBrandAndYearRange())

	tc.httpSetup.expectedHeaders = http.Header{
		"Content-Type": []string{"application/json"},
	}
	var expectedResponse interface{}
	if !tc.httpSetup.isErrorResponse {
		expectedResponse = map[string]interface{}{
			"message": tc.successMessage,
			"data":    tc.returnedVehicles,
		}
	} else {
		expectedResponse = map[string]interface{}{
			"status":  http.StatusText(tc.httpSetup.expectedStatusCode),
			"message": tc.handlerError.Error(),
		}
	}
	tc.httpSetup.expectedResponse, _ = json.Marshal(expectedResponse)

	tc.setup.mockService.On("FindByBrandAndYearRange", tc.brand, tc.startYear, tc.endYear).Return(tc.returnedVehicles, tc.serviceError)

	tc.httpSetup.req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/brand/%s/between/%v/%v", basePath, tc.brand, tc.startYear, tc.endYear), nil)
	tc.httpSetup.res = httptest.NewRecorder()
}

// Assert is a method that asserts the test case
func (tc *FindByBrandAndYearRangeTestCase) Assert(t *testing.T) {
	if tc.mockOnCalled {
		assert.True(t, tc.setup.mockService.AssertCalled(t, "FindByBrandAndYearRange", tc.brand, tc.startYear, tc.endYear))
	} else {
		assert.True(t, tc.setup.mockService.AssertNotCalled(t, "FindByBrandAndYearRange", tc.brand, tc.startYear, tc.endYear))
	}
	assert.Equal(t, tc.httpSetup.expectedStatusCode, tc.httpSetup.res.Code)
	assert.Equal(t, tc.httpSetup.expectedHeaders, tc.httpSetup.res.Header())
	assert.JSONEq(t, string(tc.httpSetup.expectedResponse), tc.httpSetup.res.Body.String())
}

func TestHandler_FindByBrandAndYearRange(t *testing.T) {
	// Create test cases
	testCases := []FindByBrandAndYearRangeTestCase{
		{
			// This test evaluates that FindByBrandAndYearRange returns 200 ok and a map of vehicles that match the brand and fabrication year range
			name:           "should return 200 ok and a map of vehicles that match the brand and fabrication year range",
			brand:          "Ford",
			startYear:      2010,
			endYear:        2012,
			successMessage: "vehicles found",
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
			mockOnCalled: true,
			httpSetup: &TestCaseHttpSetup{
				expectedStatusCode: http.StatusOK,
			},
		},
		{
			// This test evaluates that FindByBrandAndYearRange returns 400 bad request error when the start year is invalid
			name:         "should return 400 bad request error when the start year is invalid",
			brand:        "Ford",
			startYear:    "a",
			handlerError: errors.New("invalid start_year"),
			httpSetup: &TestCaseHttpSetup{
				isErrorResponse:    true,
				expectedStatusCode: http.StatusBadRequest,
			},
		},
		{
			// This test evaluates that FindByBrandAndYearRange returns 400 baad request error when the end year is invalid
			name:         "should return 400 bad request error when the end year is invalid",
			brand:        "Ford",
			startYear:    2010,
			endYear:      "a",
			handlerError: errors.New("invalid end_year"),
			httpSetup: &TestCaseHttpSetup{
				isErrorResponse:    true,
				expectedStatusCode: http.StatusBadRequest,
			},
		},
		{
			// This test evaluates that FindByBrandAndYearRange returns 500 internal server error when the service returns an error
			name:         "should return 500 internal server error when the service returns an error",
			brand:        "Ford",
			startYear:    2010,
			endYear:      2012,
			serviceError: errors.New("error"),
			handlerError: errors.New("internal error"),
			mockOnCalled: true,
			httpSetup: &TestCaseHttpSetup{
				isErrorResponse:    true,
				expectedStatusCode: http.StatusInternalServerError,
			},
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		testCase.Arrange()
		testCase.setup.server.ServeHTTP(testCase.httpSetup.res, testCase.httpSetup.req)
		testCase.Assert(t)
	}
}

// AverageMaxSpeedByBrandTestCase is a struct that represents a test case for AverageMaxSpeedByBrand
type AverageMaxSpeedByBrandTestCase struct {
	setup           *TestCaseServerSetup
	name            string
	brand           string
	successMessage  string
	returnedAverage float64
	serviceError    error
	handlerError    error
	mockOnCalled    bool
	httpSetup       *TestCaseHttpSetup
}

// Arrange is a method that sets up the test case
func (tc *AverageMaxSpeedByBrandTestCase) Arrange() {
	tc.setup = Setup()
	tc.setup.server.GET(basePath+"/average_speed/brand/:brand", tc.setup.handler.AverageMaxSpeedByBrand())

	tc.httpSetup.expectedHeaders = http.Header{
		"Content-Type": []string{"application/json"},
	}
	var expectedResponse interface{}
	if !tc.httpSetup.isErrorResponse {
		expectedResponse = map[string]interface{}{
			"message": tc.successMessage,
			"data":    tc.returnedAverage,
		}
	} else {
		expectedResponse = map[string]interface{}{
			"status":  http.StatusText(tc.httpSetup.expectedStatusCode),
			"message": tc.handlerError.Error(),
		}
	}
	tc.httpSetup.expectedResponse, _ = json.Marshal(expectedResponse)

	tc.setup.mockService.On("AverageMaxSpeedByBrand", tc.brand).Return(tc.returnedAverage, tc.serviceError)

	tc.httpSetup.req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/average_speed/brand/%s", basePath, tc.brand), nil)
	tc.httpSetup.res = httptest.NewRecorder()
}

// Assert is a method that asserts the test case
func (tc *AverageMaxSpeedByBrandTestCase) Assert(t *testing.T) {
	if tc.mockOnCalled {
		assert.True(t, tc.setup.mockService.AssertCalled(t, "AverageMaxSpeedByBrand", tc.brand))
	} else {
		assert.True(t, tc.setup.mockService.AssertNotCalled(t, "AverageMaxSpeedByBrand", tc.brand))
	}
	assert.Equal(t, tc.httpSetup.expectedStatusCode, tc.httpSetup.res.Code)
	assert.Equal(t, tc.httpSetup.expectedHeaders, tc.httpSetup.res.Header())
	assert.JSONEq(t, string(tc.httpSetup.expectedResponse), tc.httpSetup.res.Body.String())
}

// TestHandler_AverageMaxSpeedByBrand is a method that tests the AverageMaxSpeedByBrand handler
func TestHandler_AverageMaxSpeedByBrand(t *testing.T) {
	// Create test cases
	testCases := []AverageMaxSpeedByBrandTestCase{
		{
			// This test evaluates that AverageMaxSpeedByBrand returns 200 ok and the average max speed of the vehicles that match the brand
			name:            "should return 200 ok and the average max speed of the vehicles that match the brand",
			brand:           "Ford",
			successMessage:  "average max speed found",
			returnedAverage: 180,
			mockOnCalled:    true,
			httpSetup: &TestCaseHttpSetup{
				expectedStatusCode: http.StatusOK,
			},
		},
		{
			// This test evaluates that AverageMaxSpeedByBrand returns 404 not found error when there are no vehicles that match the brand
			name:         "should return 404 not found error when there are no vehicles that match the brand",
			brand:        "Chevrolet",
			serviceError: internal.ErrServiceNoVehicles,
			handlerError: errors.New("vehicles not found"),
			mockOnCalled: true,
			httpSetup: &TestCaseHttpSetup{
				isErrorResponse:    true,
				expectedStatusCode: http.StatusNotFound,
			},
		},
		{
			// This test evaluates that AverageMaxSpeedByBrand returns 500 internal server error when the service returns an error
			name:         "should return 500 internal server error when the service returns an error",
			brand:        "Ford",
			serviceError: errors.New("error"),
			handlerError: errors.New("internal error"),
			mockOnCalled: true,
			httpSetup: &TestCaseHttpSetup{
				isErrorResponse:    true,
				expectedStatusCode: http.StatusInternalServerError,
			},
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		testCase.Arrange()
		testCase.setup.server.ServeHTTP(testCase.httpSetup.res, testCase.httpSetup.req)
		testCase.Assert(t)
	}
}

// AverageCapacityByBrandTestCase is a struct that represents a test case for AverageCapacityByBrand
type AverageCapacityByBrandTestCase struct {
	setup            *TestCaseServerSetup
	name             string
	brand            string
	successMessage   string
	returnedCapacity int
	serviceError     error
	handlerError     error
	mockOnCalled     bool
	httpSetup        *TestCaseHttpSetup
}

// Arrange is a method that sets up the test case
func (tc *AverageCapacityByBrandTestCase) Arrange() {
	tc.setup = Setup()
	tc.setup.server.GET(basePath+"/average_capacity/brand/:brand", tc.setup.handler.AverageCapacityByBrand())

	tc.httpSetup.expectedHeaders = http.Header{
		"Content-Type": []string{"application/json"},
	}
	var expectedResponse interface{}
	if !tc.httpSetup.isErrorResponse {
		expectedResponse = map[string]interface{}{
			"message": tc.successMessage,
			"data":    tc.returnedCapacity,
		}
	} else {
		expectedResponse = map[string]interface{}{
			"status":  http.StatusText(tc.httpSetup.expectedStatusCode),
			"message": tc.handlerError.Error(),
		}
	}
	tc.httpSetup.expectedResponse, _ = json.Marshal(expectedResponse)

	tc.setup.mockService.On("AverageCapacityByBrand", tc.brand).Return(tc.returnedCapacity, tc.serviceError)

	tc.httpSetup.req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/average_capacity/brand/%s", basePath, tc.brand), nil)
	tc.httpSetup.res = httptest.NewRecorder()
}

// Assert is a method that asserts the test case
func (tc *AverageCapacityByBrandTestCase) Assert(t *testing.T) {
	if tc.mockOnCalled {
		assert.True(t, tc.setup.mockService.AssertCalled(t, "AverageCapacityByBrand", tc.brand))
	} else {
		assert.True(t, tc.setup.mockService.AssertNotCalled(t, "AverageCapacityByBrand", tc.brand))
	}
	assert.Equal(t, tc.httpSetup.expectedStatusCode, tc.httpSetup.res.Code)
	assert.Equal(t, tc.httpSetup.expectedHeaders, tc.httpSetup.res.Header())
	assert.JSONEq(t, string(tc.httpSetup.expectedResponse), tc.httpSetup.res.Body.String())
}

// TestHandler_AverageCapacityByBrand is a method that tests the AverageCapacityByBrand handler
func TestHandler_AverageCapacityByBrand(t *testing.T) {
	// Create test cases
	testCases := []AverageCapacityByBrandTestCase{
		{
			// This test evaluates that AverageCapacityByBrand returns 200 ok and the average capacity of the vehicles that match the brand
			name:             "should return 200 ok and the average capacity of the vehicles that match the brand",
			brand:            "Ford",
			successMessage:   "average capacity found",
			returnedCapacity: 5,
			mockOnCalled:     true,
			httpSetup: &TestCaseHttpSetup{
				expectedStatusCode: http.StatusOK,
			},
		},
		{
			// This test evaluates that AverageCapacityByBrand returns 404 not found error when there are no vehicles that match the brand
			name:         "should return 404 not found error when there are no vehicles that match the brand",
			brand:        "Chevrolet",
			serviceError: internal.ErrServiceNoVehicles,
			handlerError: errors.New("vehicles not found"),
			mockOnCalled: true,
			httpSetup: &TestCaseHttpSetup{
				isErrorResponse:    true,
				expectedStatusCode: http.StatusNotFound,
			},
		},
		{
			// This test evaluates that AverageCapacityByBrand returns 500 internal server error when the service returns an error
			name:         "should return 500 internal server error when the service returns an error",
			brand:        "Ford",
			serviceError: errors.New("error"),
			handlerError: errors.New("internal error"),
			mockOnCalled: true,
			httpSetup: &TestCaseHttpSetup{
				isErrorResponse:    true,
				expectedStatusCode: http.StatusInternalServerError,
			},
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		testCase.Arrange()
		testCase.setup.server.ServeHTTP(testCase.httpSetup.res, testCase.httpSetup.req)
		testCase.Assert(t)
	}
}

// SearchByWeightRangeTestCase is a struct that represents a test case for SearchByWeightRange
type SearchByWeightRangeTestCase struct {
	setup            *TestCaseServerSetup
	name             string
	fromWeight       interface{}
	toWeight         interface{}
	ok               bool
	successMessage   string
	returnedVehicles map[int]internal.Vehicle
	serviceError     error
	handlerError     error
	mockOnCalled     bool
	httpSetup        *TestCaseHttpSetup
}

// Arrange is a method that sets up the test case
func (tc *SearchByWeightRangeTestCase) Arrange() {
	tc.setup = Setup()
	tc.setup.server.GET(basePath+"/weight", tc.setup.handler.SearchByWeightRange())

	tc.httpSetup.expectedHeaders = http.Header{
		"Content-Type": []string{"application/json"},
	}
	var expectedResponse interface{}
	if !tc.httpSetup.isErrorResponse {
		expectedResponse = map[string]interface{}{
			"message": tc.successMessage,
			"data":    tc.returnedVehicles,
		}
	} else {
		expectedResponse = map[string]interface{}{
			"status":  http.StatusText(tc.httpSetup.expectedStatusCode),
			"message": tc.handlerError.Error(),
		}
	}
	tc.httpSetup.expectedResponse, _ = json.Marshal(expectedResponse)

	var query internal.SearchQuery
	var target string

	if tc.ok {
		if tc.fromWeight == nil {
			query = internal.SearchQuery{
				FromWeight: 0.0,
				ToWeight:   tc.toWeight.(float64),
			}
			target = fmt.Sprintf("%s/weight?weight_min=%v", basePath, tc.fromWeight)
		} else if tc.toWeight == nil {
			query = internal.SearchQuery{
				FromWeight: tc.fromWeight.(float64),
				ToWeight:   0.0,
			}
			target = fmt.Sprintf("%s/weight?weight_max=%v", basePath, tc.toWeight)
		} else {
			query = internal.SearchQuery{
				FromWeight: tc.fromWeight.(float64),
				ToWeight:   tc.toWeight.(float64),
			}
			target = fmt.Sprintf("%s/weight?weight_min=%v&weight_max=%v", basePath, tc.fromWeight, tc.toWeight)
		}
	} else {
		target = fmt.Sprintf("%s/weight", basePath)
	}

	tc.setup.mockService.On("SearchByWeightRange", query, tc.ok).Return(tc.returnedVehicles, tc.serviceError)

	tc.httpSetup.req = httptest.NewRequest(http.MethodGet, target, nil)
	tc.httpSetup.res = httptest.NewRecorder()
}

// Assert is a method that asserts the test case
func (tc *SearchByWeightRangeTestCase) Assert(t *testing.T) {
	if tc.mockOnCalled {
		query := internal.SearchQuery{
			FromWeight: tc.fromWeight.(float64),
			ToWeight:   tc.toWeight.(float64),
		}
		assert.True(t, tc.setup.mockService.AssertCalled(t, "SearchByWeightRange", query, tc.ok))
	} else {
		var query internal.SearchQuery
		assert.True(t, tc.setup.mockService.AssertNotCalled(t, "SearchByWeightRange", query, tc.ok))
	}
	assert.Equal(t, tc.httpSetup.expectedStatusCode, tc.httpSetup.res.Code)
	assert.Equal(t, tc.httpSetup.expectedHeaders, tc.httpSetup.res.Header())
	assert.JSONEq(t, string(tc.httpSetup.expectedResponse), tc.httpSetup.res.Body.String())
}

// TestHandler_SearchByWeightRange is a method that tests the SearchByWeightRange handler
func TestHandler_SearchByWeightRange(t *testing.T) {
	// Create test cases
	testCases := []SearchByWeightRangeTestCase{
		{
			// This test evaluates that SearchByWeightRange returns 200 ok and the vehicles that match the weight range
			name:           "should return 200 ok and the vehicles that match the weight range",
			fromWeight:     1000.0,
			toWeight:       1050.0,
			ok:             true,
			successMessage: "vehicles found",
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
			},
			mockOnCalled: true,
			httpSetup: &TestCaseHttpSetup{
				expectedStatusCode: http.StatusOK,
			},
		},
		{
			// This test evaluates that SearchByWeightRange returns 500 internal server error when the service returns an error
			name:         "should return 500 internal server error when the service returns an error",
			fromWeight:   1000.0,
			toWeight:     1050.0,
			ok:           true,
			serviceError: errors.New("error"),
			handlerError: errors.New("internal error"),
			mockOnCalled: true,
			httpSetup: &TestCaseHttpSetup{
				isErrorResponse:    true,
				expectedStatusCode: http.StatusInternalServerError,
			},
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		testCase.Arrange()
		testCase.setup.server.ServeHTTP(testCase.httpSetup.res, testCase.httpSetup.req)
		testCase.Assert(t)
	}
}
