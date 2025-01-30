package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.60

import (
	"context"
	"fmt"
	"log"

	"github.com/juliya711/gql-yt/database"
	"github.com/juliya711/gql-yt/graph/model"
)

// CreateJobListing is the resolver for the createJobListing field.
func (r *mutationResolver) CreateJobListing(ctx context.Context, input model.CreateJobListingInput) (*model.JobListing, error) {
	db := database.Connect()
	return db.CreateJobListing(input), nil
}

// UpdateJobListing is the resolver for the updateJobListing field.
func (r *mutationResolver) UpdateJobListing(ctx context.Context, id string, input model.UpdateJobListingInput) (*model.JobListing, error) {
	db := database.Connect()
	return db.UpdateJobListing(id, input), nil
}

// DeleteJobListing is the resolver for the deleteJobListing field.
func (r *mutationResolver) DeleteJobListing(ctx context.Context, id string) (*model.DeleteJobResponse, error) {
	db := database.Connect()
	return db.DeleteJobListing(id), nil
}

// Jobs is the resolver for the jobs field.
func (r *queryResolver) Jobs(ctx context.Context) ([]*model.JobListing, error) {
	db := database.Connect()
	return db.GetJobs(), nil
}

// Job is the resolver for the job field.
func (r *queryResolver) Job(ctx context.Context, id string) (*model.JobListing, error) {
	db := database.Connect()

	return db.GetJob(id), nil
}

// ExportDevicesDiscovered is the resolver for the exportDevicesDiscovered field.
func (r *queryResolver) ExportDevicesDiscovered(ctx context.Context, input model.ExportDevicesDiscoveredInput) (*model.DevicesDiscoveredResponse, error) {
	devices, err := database.FetchDiscoveredDevices(input.CompanyID, input.AssessmentID)
	if err != nil {
		log.Printf("Error fetching discovered devices: %v", err)
		return nil, fmt.Errorf("failed to fetch discovered devices: %v", err)
	}

	var discoveredDevices []*model.DeviceDiscovered
	for _, device := range devices {
		discoveredDevices = append(discoveredDevices, &model.DeviceDiscovered{
			ProductIdentifier: interfaceToStringPtr(device["product_identifier"]),
			Hostname:          interfaceToStringPtr(device["hostname"]),
			Vendor:            interfaceToStringPtr(device["vendor_cd"]),
			SerialNumber:      interfaceToStringPtr(device["serialNumber"]),
			IPAddresses:       interfaceToStringPtr(device["IP_Address"]),
		})
	}

	return &model.DevicesDiscoveredResponse{
		DevicesDiscoveredSuccessfully: discoveredDevices,
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// Helper function to convert interface{} to *string
func interfaceToStringPtr(value interface{}) *string {
	if value == nil {
		return nil
	}
	strValue, ok := value.(string)
	if !ok {
		return nil
	}
	return &strValue
}
