
# Start app container
start-app:
	docker-compose up -d web

# Start db container
start-db:
	docker-compose up -d db

# Run all containers listed in the docker-compose file
start-all:
	docker-compose up -d

# Format code to Go format standards
format:
	gofmt -w .

# Verify if the code match the golang standards
validate-conventions:
	golint ./...

# Verify codes that can run into problems
validate-code:
	go vet ./...

# Run all tests
test:
	go test./...

# Run all tests and generate coverage file.
# The coverage will be displayed in the browser
test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
