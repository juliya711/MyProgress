# Quickstart GoLang GraphQL API

This is a sample GoLang-based GraphQL API for the HWI MongoDB custom-hosted solution. The project demonstrates how to build a GraphQL API using gqlgen, integrate with MongoDB, containerize with Docker, and deploy using Kubernetes.

---

## What's Included?

This repository includes:

- _GoLang & gqlgen:_  
  A GraphQL API built in GoLang using [gqlgen](https://gqlgen.com/) for code generation.

- _MongoDB Integration:_  
  CRUD operations and advanced aggregation pipelines using the official [mongo-go-driver](https://pkg.go.dev/go.mongodb.org/mongo-driver).

- _Docker:_  
  A Dockerfile for containerizing the application for development and production environments.

- _Kubernetes:_  
  Configuration files for deploying and scaling the application on Kubernetes clusters.

- _Advanced MongoDB Operations:_  
  Functions to fetch and aggregate device data for integration with the ATNA dashboard.

---

## Prerequisites

1. _Go:_  
   Install Go (version 1.18 or later).

2. _MongoDB:_  
   A MongoDB instance (either custom-hosted or via MongoDB Atlas).

3. _Docker:_  
   For containerizing the application.

4. _Kubernetes:_  
   (Optional) For orchestrating deployments in production.

5. _gqlgen (optional):_  
   Install gqlgen globally if desired:
   ```bash
   go install github.com/99designs/gqlgen@latest
   ```

Development

1. Clone and Setup

Clone the repository and navigate into the project directory:

git clone <REPO_URL>
cd gql-yt

2. Install Dependencies

Ensure Go modules are up to date:

go mod tidy

3. Generate GraphQL Code

Run gqlgen to generate the necessary GraphQL code:

go run github.com/99designs/gqlgen

4. Run the Application

Start the GraphQL server by running:

go run server.go

Your GraphQL endpoint will be available at http://localhost:8080.

Project Structure

gql-yt/
├── database/  
│ └── database.go # MongoDB connection & CRUD operations.
├── graph/  
│ ├── generated/ # Auto-generated code by gqlgen.
│ ├── model/ # Go structures mapping to GraphQL types.
│ ├── schema.graphqls # GraphQL schema definitions.
│ ├── schema.resolvers.go # Auto-generated resolver stubs.
│ ├── resolver.go # Custom resolver implementations.
│ └── resolver_test.go # Unit tests for resolver functions.
├── go.mod # Go module dependencies.
├── go.sum # Go module checksums.
├── gqlgen.yml # gqlgen configuration file.
├── server.go # Entry point of the GraphQL server.
└── tools.go # Developer tools for code generation.

Docker & Kubernetes Deployment

Docker

Dockerfile

Below is an example Dockerfile to containerize the application:

FROM golang:1.18-alpine

WORKDIR /app

# Copy and download dependencies

COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build the application

COPY . .
RUN go build -o server .

# Expose the application port

EXPOSE 8080

# Run the executable

CMD ["./server"]

Build and Run the Docker Container

Build the Docker image:

docker build -t gql-yt .

Run the Docker container:

docker run -p 8080:8080 gql-yt

Kubernetes

Deploy the application using Kubernetes configuration files. An example deployment.yaml might look like:

apiVersion: apps/v1
kind: Deployment
metadata:
name: gql-yt-deployment
spec:
replicas: 3
selector:
matchLabels:
app: gql-yt
template:
metadata:
labels:
app: gql-yt
spec:
containers: - name: gql-yt
image: gql-yt:latest
ports: - containerPort: 8080

Apply your Kubernetes configurations:

kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
kubectl apply -f ingress.yaml

MongoDB Operations

The project integrates with MongoDB using the mongo-go-driver. Key functions include:
• Connect:
Establish a connection to MongoDB.
• Aggregation Pipeline:
• FetchDiscoveredDevices: Execute an aggregation pipeline to fetch device data for the ATNA dashboard integration.

GraphQL Schema

The GraphQL schema is defined in graph/schema.graphqls and includes the following:

Types:
• DeviceDiscovered
• DevicesDiscoveredResponse

Queries:
• exportDevicesDiscovered(input: ExportDevicesDiscoveredInput!): DevicesDiscoveredResponse!: Fetch discovered devices based on company and assessment IDs.

Releasing & Contributing
• Contributing:
Contributions are welcome! Please open an issue or submit a pull request with improvements or bug fixes.
• Releasing:
Use standard Git versioning practices. Tag pushes can trigger Docker builds and Kubernetes deployments as part of your CI/CD pipeline.

License

This project is licensed under the MIT License.

Additional Resources
• gqlgen Documentation
• MongoDB Go Driver Documentation
• Docker Documentation
• Kubernetes Documentation

---
