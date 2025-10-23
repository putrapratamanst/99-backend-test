@echo off
setlocal enabledelayedexpansion

if "%1"=="" goto help
if "%1"=="help" goto help
if "%1"=="generate-all" goto generate-all
if "%1"=="run-all" goto run-all
goto help

:help
echo Available commands:
echo   run.bat generate-all  - Generate and build all services (recommended)
echo   run.bat run-all       - Run all services simultaneously
goto end

:install-deps
echo Installing Go dependencies...
go mod tidy
go mod download
echo Dependencies installed!
goto end

:generate-all
echo Creating bin directory...
if not exist bin mkdir bin
echo Installing dependencies...
call :install-deps
echo Installing Python dependencies...
pip install -r python-libs.txt
echo Setting CGO_ENABLED=0 for pure Go build...
set CGO_ENABLED=0
echo Building user service...
go build -o bin/user-service.exe ./cmd/user-service
if errorlevel 1 (
    echo Error building user service
    goto end
)
echo Building public API...
go build -o bin/public-api.exe ./cmd/public-api
if errorlevel 1 (
    echo Error building public API
    goto end
)
echo Python listing service ready to run...
echo All services built successfully!
goto end

:run-all
call :generate-all
if errorlevel 1 goto end
echo Starting all services...
echo Starting user service on port 8001...
start /min "User Service" bin/user-service.exe
timeout /t 2 /nobreak >nul
echo Starting listing service (Python) on port 6000...
start /min "Listing Service" python listing_service.py
timeout /t 2 /nobreak >nul
echo Starting public API on port 8000...
start /min "Public API" bin/public-api.exe
echo All services are starting...
echo Check Task Manager or visit:
echo   - Public API: http://localhost:8000
echo   - User Service: http://localhost:8001
echo   - Listing Service: http://localhost:6000/listings/ping
goto end

:end
endlocal