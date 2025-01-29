package database

import (
	"context"
	"log"
	"time"

	"github.com/juliya711/gql-yt/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connectionString string = "mongodb+srv://Juliya:jp12345@cluster0.gmcyl.mongodb.net/?retryWrites=true&w=majority"

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		client: client,
	}
}

func (db *DB) GetJob(id string) *model.JobListing {
	jobCollec := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var jobListing model.JobListing
	err := jobCollec.FindOne(ctx, filter).Decode(&jobListing)
	if err != nil {
		log.Fatal(err)
	}
	return &jobListing
}

func (db *DB) GetJobs() []*model.JobListing {
	jobCollec := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var jobListings []*model.JobListing
	cursor, err := jobCollec.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), &jobListings); err != nil {
		panic(err)
	}

	return jobListings
}

func (db *DB) CreateJobListing(jobInfo model.CreateJobListingInput) *model.JobListing {
	jobCollec := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	inserg, err := jobCollec.InsertOne(ctx, bson.M{"title": jobInfo.Title, "description": jobInfo.Description, "url": jobInfo.URL, "company": jobInfo.Company})

	if err != nil {
		log.Fatal(err)
	}

	insertedID := inserg.InsertedID.(primitive.ObjectID).Hex()
	returnJobListing := model.JobListing{ID: insertedID, Title: jobInfo.Title, Company: jobInfo.Company, Description: jobInfo.Description, URL: jobInfo.URL}
	return &returnJobListing
}

func (db *DB) UpdateJobListing(jobId string, jobInfo model.UpdateJobListingInput) *model.JobListing {
	jobCollec := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	updateJobInfo := bson.M{}

	// if jobInfo.Title != nil {
	// 	updateJobInfo["title"] = jobInfo.Title
	// }
	// if jobInfo.Description != nil {
	// 	updateJobInfo["description"] = jobInfo.Description
	// }
	// if jobInfo.URL != nil {
	// 	updateJobInfo["url"] = jobInfo.URL
	// }

	_id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateJobInfo}

	results := jobCollec.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var jobListing model.JobListing

	if err := results.Decode(&jobListing); err != nil {
		log.Fatal(err)
	}

	return &jobListing
}

func (db *DB) DeleteJobListing(jobId string) *model.DeleteJobResponse {
	jobCollec := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": _id}
	_, err := jobCollec.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return &model.DeleteJobResponse{DeleteJobID: jobId}

	// GetDiscoveredDevices fetches devices for the given company_id from the assets collection.
}

// GetDiscoveredDevices fetches devices for the given company_id from the assets collection.
func FetchDiscoveredDevices(companyID, assessmentID string) ([]map[string]interface{}, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Juliya:jp12345@cluster0.gmcyl.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("ATNA").Collection(`assets`)

	// Aggregation pipeline
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$addFields", Value: bson.D{
			{Key: "ip_addresses", Value: bson.D{
				{Key: "$arrayElemAt", Value: bson.A{"$ip_addresses", 0}},
			}},
		}}},
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "software"},
			{Key: "let", Value: bson.D{
				{Key: "d_company_id", Value: "$company_id"},
				{Key: "d_assessment_id", Value: "$assessment_id"},
				{Key: "d_serialNumber", Value: "$serial_number"},
			}},
			{Key: "pipeline", Value: bson.A{
				bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{
					{Key: "$and", Value: bson.A{
						bson.D{{Key: "$eq", Value: bson.A{"$company_id", "$$d_company_id"}}},
						bson.D{{Key: "$eq", Value: bson.A{"$assessment_id", "$$d_assessment_id"}}},
						bson.D{{Key: "$eq", Value: bson.A{"$device_serial_number", "$$d_serialNumber"}}},
					}},
				}}}}},
			}},
			{Key: "as", Value: "softwareLookup"},
		}}},
		bson.D{{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$softwareLookup"},
			{Key: "preserveNullAndEmptyArrays", Value: true},
		}}},
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "module"},
			{Key: "localField", Value: "g_module_id.id"},
			{Key: "foreignField", Value: "module_id"},
			{Key: "as", Value: "moduleLookup"},
		}}},
		bson.D{{Key: "$addFields", Value: bson.D{
			{Key: "moduleLookup", Value: bson.D{
				{Key: "$arrayElemAt", Value: bson.A{"$moduleLookup", 0}},
			}},
		}}},
		bson.D{{Key: "$match", Value: bson.D{
			{Key: "moduleLookup.type", Value: bson.D{{Key: "$ne", Value: "PORT"}}},
		}}},
		bson.D{{Key: "$project", Value: bson.D{
			{Key: "product_identifier", Value: 1},
			{Key: "hostname", Value: bson.D{{Key: "$toLower", Value: "$hostname"}}},
			{Key: "vendor_cd", Value: 1},
			{Key: "department", Value: "$ntt_itsm_department"},
			{Key: "replacementProduct", Value: "$migration_pid"},
			{Key: "contractedStatus", Value: "$ntt_contracted_status"},
			{Key: "country", Value: "$ntt_itsm_country"},
			{Key: "locations", Value: "$ntt_itsm_location"},
			{Key: "serialNumber", Value: "$serial_number"},
			{Key: "IP_Address", Value: "$ip_addresses"},
			{Key: "end_of_sale", Value: bson.D{
				{Key: "$dateToString", Value: bson.D{
					{Key: "format", Value: "%Y-%m-%d"},
					{Key: "date", Value: "$end_of_sale"},
				}},
			}},
			{Key: "end_of_support", Value: bson.D{
				{Key: "$dateToString", Value: bson.D{
					{Key: "format", Value: "%Y-%m-%d"},
					{Key: "date", Value: "$last_day_of_support"},
				}},
			}},
			{Key: "product_family", Value: 1},
			{Key: "product_category", Value: 1},
			{Key: "software_type", Value: "$softwareLookup.os_name"},
			{Key: "software_version", Value: "$softwareLookup.os_version"},
			{Key: "module_serial_number", Value: "$moduleLookup.serial_number"},
			{Key: "type", Value: 1},
		}}},
	}
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}
func (db *DB) Client() *mongo.Client {
	return db.client
}
