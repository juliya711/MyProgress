package graph

import (
	"context"
	"errors"
	"testing"

	"github.com/juliya711/gql-yt/graph/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking the database function
type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) FetchDiscoveredDevices(companyID, assessmentID string) ([]map[string]interface{}, error) {
	args := m.Called(companyID, assessmentID)
	result, _ := args.Get(0).([]map[string]interface{})
	return result, args.Error(1)
}

func TestExportDevicesDiscovered(t *testing.T) {
	mockDB := new(MockDatabase)
	resolver := &Resolver{}

	t.Run("Success Case", func(t *testing.T) {
		mockDevices := []map[string]interface{}{
			{
				"product_identifier": "MR46-HW",
				"hostname":           "mrpg-apparel-g-ap12",
				"vendor":             nil,
				"serial_number":      "",
				"ip_addresses":       "10.1.181.44",
			},
		}

		// Mock FetchDiscoveredDevices response
		mockDB.On("FetchDiscoveredDevices", "company1", "assessment1").Return(mockDevices, nil)

		// Call resolver function
		input := model.ExportDevicesDiscoveredInput{CompanyID: "company1", AssessmentID: "assessment1"}
		result, err := resolver.ExportDevicesDiscovered(context.Background(), &input)

		// Assertions
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.DevicesDiscoveredSuccessfully, 3)
		assert.Equal(t, "MR46-HW", *result.DevicesDiscoveredSuccessfully[0].ProductIdentifier)
		assert.Equal(t, "mrpg-apparel-g-ap12", *result.DevicesDiscoveredSuccessfully[0].Hostname)
		assert.Nil(t, result.DevicesDiscoveredSuccessfully[0].Vendor)
		assert.Equal(t, "", *result.DevicesDiscoveredSuccessfully[0].SerialNumber)
		assert.Equal(t, "10.1.181.44", *result.DevicesDiscoveredSuccessfully[0].IPAddresses)

		mockDB.AssertExpectations(t)
	})

	t.Run("Database Error", func(t *testing.T) {
		mockDB.On("FetchDiscoveredDevices", "company3", "assessment3").Return(nil, errors.New("database error"))

		input := model.ExportDevicesDiscoveredInput{CompanyID: "company3", AssessmentID: "assessment3"}
		result, err := resolver.ExportDevicesDiscovered(context.Background(), &input)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "database error")

		mockDB.AssertExpectations(t)
	})

	t.Run("No Devices Found", func(t *testing.T) {
		mockDB.On("FetchDiscoveredDevices", "company4", "assessment4").Return([]map[string]interface{}{}, nil)

		input := model.ExportDevicesDiscoveredInput{CompanyID: "company4", AssessmentID: "assessment4"}
		result, err := resolver.ExportDevicesDiscovered(context.Background(), &input)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.DevicesDiscoveredSuccessfully, 0)

		mockDB.AssertExpectations(t)
	})

	t.Run("Missing Fields", func(t *testing.T) {
		mockDevices := []map[string]interface{}{
			{
				"product_identifier": nil,
				"hostname":           "unknown-device",
				"vendor":             nil,
				"serial_number":      nil,
				"ip_addresses":       "",
			},
		}

		mockDB.On("FetchDiscoveredDevices", "company5", "assessment5").Return(mockDevices, nil)

		input := model.ExportDevicesDiscoveredInput{CompanyID: "company5", AssessmentID: "assessment5"}
		result, err := resolver.ExportDevicesDiscovered(context.Background(), &input)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.DevicesDiscoveredSuccessfully, 1)
		assert.Nil(t, result.DevicesDiscoveredSuccessfully[0].ProductIdentifier)
		assert.Equal(t, "unknown-device", *result.DevicesDiscoveredSuccessfully[0].Hostname)
		assert.Nil(t, result.DevicesDiscoveredSuccessfully[0].Vendor)
		assert.Nil(t, result.DevicesDiscoveredSuccessfully[0].SerialNumber)
		assert.Equal(t, "", *result.DevicesDiscoveredSuccessfully[0].IPAddresses)

		mockDB.AssertExpectations(t)
	})
}
