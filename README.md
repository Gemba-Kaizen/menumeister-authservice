# example-microservice-gRPC
This repositroy shows an example of how you could set up a microservice that has a gRPC server and accepts gRPC requests from clients. It uses gorm ORM to communicate with the database to create and get records.

## Third party libraries
1. gorm
2. postgres

## Start db
make db
docker exec -it database psql -U pg -W crud
- U: user
- W: flag to say password needed
- crud: database name

## Architecture
- repository
Data access layer
- services
Business logic layer
- handler
API layer