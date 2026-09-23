@echo off
echo ========================================================
echo   Starting TEMENIN AI Companion Backend Service
echo ========================================================

where go >nul 2>nul
if %errorlevel% equ 0 (
    echo [OK] Go compiler detected. Starting Go server natively...
    cd %~dp0
    go run ./cmd/server/main.go
    goto end
)

where docker >nul 2>nul
if %errorlevel% equ 0 (
    echo [OK] Docker detected. Starting backend via Docker Compose...
    cd %~dp0\..
    docker-compose up backend
    goto end
)

echo [!] Notice: Go compiler or Docker is not installed in current PATH.
echo [!] The Mobile App includes built-in mock services with local storage fallback
echo [!] You can install Go via: winget install GoLang.Go
echo ========================================================
:end
pause
