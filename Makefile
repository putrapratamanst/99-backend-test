.PHONY: generate-all run-all

help:
	@echo "Available commands:"
	@echo "  generate-all  - Generate and build all services (recommended)"
	@echo "  run-all       - Run all services simultaneously"

install-deps:
	go mod tidy
	go mod download

generate-all: install-deps
	@echo "Creating bin directory..."
	@if not exist bin mkdir bin
	@echo "Installing Python dependencies..."
	pip install -r python-libs.txt
	@echo "Building user service..."
	set CGO_ENABLED=0 && go build -o bin/user-service.exe ./cmd/user-service
	@echo "Building public API..."
	set CGO_ENABLED=0 && go build -o bin/public-api.exe ./cmd/public-api
	@echo "Python listing service ready to run..."
	@echo "All services built successfully!"

run-all: generate-all
	@echo "Starting all services..."
	@echo "Starting user service on port 8001..."
	start /min powershell -Command "bin/user-service.exe"
	@timeout /t 2 /nobreak >nul
	@echo "Starting listing service (Python) on port 6000..."
	start /min powershell -Command "python listing_service.py"
	@timeout /t 2 /nobreak >nul
	@echo "Starting public API on port 8000..."
	start /min powershell -Command "bin/public-api.exe"
	@echo "All services are starting... Check task manager or ports 8000, 8001, 6000"