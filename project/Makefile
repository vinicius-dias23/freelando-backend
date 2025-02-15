AUTH_BINARY=authApp
CHAT_BINARY=chatApp
CORE_BINARY=coreApp
PROFILE_BINARY=profileApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker compose up -d
	@echo "Docker images started!"

## up_build: stops docker compose (if running), builds all projects and starts docker compose
up_build: build_auth build_chat build_core build_profile
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Building (when required) and starting docker images..."
	docker compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker compose down
	@echo "Done!"

## build_auth: builds the auth binary as a linux executable
build_auth:
	@echo "Building auth binary..."
	cd ../authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./api/cmd
	@echo "Done!"

## build_chat: builds the chat binary as a linux executable
build_chat:
	@echo "Building chat binary..."
	cd ../chat-service && env GOOS=linux CGO_ENABLED=0 go build -o ${CHAT_BINARY} .
	@echo "Done!"

## build_core: builds the core binary as a linux executable
build_core:
	@echo "Building core binary..."
	cd ../core-service && env GOOS=linux CGO_ENABLED=0 go build -o ${CORE_BINARY} .
	@echo "Done!"

## build_profile: builds the profile binary as a linux executable
build_profile:
	@echo "Building profile binary..."
	cd ../profile-service && env GOOS=linux CGO_ENABLED=0 go build -o ${PROFILE_BINARY} .
	@echo "Done!"