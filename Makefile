# Include values from .env file
# include .env

# ================================================================================
# HELPERS
# ================================================================================

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

# Confirmation target
.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N]' && read ans && [ $${ans:-N} = y ]



# ================================================================================
# DEVELOPMENT
# ================================================================================

## run: run the ./cmd/api/ application
.PHONY: run
run:
	go run ./cmd/api


## db/migrations/new name=$1: creates a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}'
	goose -dir=./internal/schema create ${name} sql

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	goose -dir=./internal/schema sqlite ./spotfinder.db up

## db/migrations/down: apply all down database migrations
.PHONY: db/migrations/down
db/migrations/down:
	@echo 'Running down migrations...'
	goose -dir=./internal/schema sqlite ./spotfinder.db down



# ================================================================================
# QUALITY CONTROL
# ================================================================================

## audit: tidy dependencies and forma, vet and test all code
.PHONY: audit
audit:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify

	@echo 'Formatting code...'
	go fmt ./...

	@echo 'Vetting code...'
	go vet ./...

	@echo 'Running tests...'
	go test ./... -v -count=1



# ================================================================================
# BUILD
# ================================================================================

## build: build the cmd/api application
.PHONY: build
build:
	@echo 'Building cmd/api...'
	go build -ldflags="-s" -o=./bin/spotfinder ./cmd/api
	@echo 'Building for linux...'
	GOOS=linux GOARCH=amd64 go build -ldflags="-s" -o=./bin/spotfinder-linux-amd64 ./cmd/api
	@echo 'Building for OSX/darwin...'
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s" -o=./bin/spotfinder-darwin-amd64 ./cmd/api
	@echo 'Building for windows...'
	GOOS=windows GOARCH=amd64 go build -ldflags="-s" -o=./bin/spotfinder-windows-amd64.exe ./cmd/api



# ================================================================================
# CLEAN
# ================================================================================

## clean: clean up files
.PHONY: clean
clean:
	@echo 'Cleaning up files...'
	rm -rvf ./bin/*
	@echo 'Cleaning up folders...'
	rm -rvf ./bin/
	@echo 'Cleaning database...'
	rm -rfv ./*.db
