# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/kurses: run the cmd/kurses application
.PHONY: run/kurses
run/kurses:
	go run ./cmd/kurses

## audit: tidy and vendor dependencies and format, vet and test all code
.PHONY: audit
audit: vendor
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

# ==================================================================================== #
# BUILD
# ==================================================================================== #
## build/api: build the cmd/api application
.PHONY: build/kurses
build/kurses:
	@echo 'Building cmd/kurses...'
	go build -o=./bin/kurses ./cmd/kurses
#	go build -ldflags='-s -w' -o=./bin/api ./cmd/api
#	GOOS=linux GOARCH=amd64 go build -o=./bin/linux_amd64/kurses ./cmd/kurses
# strip dev
#	GOOS=linux GOARCH=amd64 go build -ldflags='-s -w' -o=./bin/linux_amd64/api ./cmd/api
