package graph

import (
	"context"
	"log"

	"github.com/juliya711/gql-yt/database"
	"github.com/juliya711/gql-yt/graph/model"
)

// Helper function to convert string to *string
func toPointer(value string) *string {
	return &value
}

type Resolver struct{}

func (r *Resolver) ExportDevicesDiscovered(ctx context.Context, input *model.ExportDevicesDiscoveredInput) (*model.DevicesDiscoveredResponse, error) {
	devices, err := database.FetchDiscoveredDevices(input.CompanyID, input.AssessmentID)
	if err != nil {
		log.Printf("Error fetching devices: %v", err)
		return nil, err
	}

	var discoveredDevices []*model.DeviceDiscovered
	for _, device := range devices {
		productIdentifier, _ := device["product_identifier"].(string)
		hostname, _ := device["hostname"].(string)
		vendor, _ := device["vendor"].(string)
		serialNumber, _ := device["serial_number"].(string)
		ipAddresses, _ := device["ip_addresses"].(string)

		discoveredDevices = append(discoveredDevices, &model.DeviceDiscovered{
			ProductIdentifier: toPointer(productIdentifier),
			Hostname:          toPointer(hostname),
			Vendor:            toPointer(vendor),
			SerialNumber:      toPointer(serialNumber),
			IPAddresses:       toPointer(ipAddresses),
		})
	}

	return &model.DevicesDiscoveredResponse{
		DevicesDiscoveredSuccessfully: discoveredDevices,
	}, nil
}
