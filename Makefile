tidy: 
	go mod tidy
	go mod vendor

install-tools:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.3.0

generate gen: 
	$(GOPATH)/bin/oapi-codegen -generate fiber -package=api_gen -o ./adapter/driving_adapter/rest_driving_adapter/gen/fiber.go ./adapter/driving_adapter/rest_driving_adapter/gen/api_spec.yml 
	$(GOPATH)/bin/oapi-codegen -generate types -package=api_gen -o ./adapter/driving_adapter/rest_driving_adapter/gen/types.go ./adapter/driving_adapter/rest_driving_adapter/gen/api_spec.yml 

run:
	go run main.go

compose:
	docker-compose up