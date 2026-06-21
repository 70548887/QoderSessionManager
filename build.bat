@echo off
echo Building QoderSessionManager for Windows...

set VERSION=1.0.0
set LDFLAGS=-X "qoder-sm/pkg/qoder.Version=%VERSION%"

if not exist build\bin mkdir build\bin
echo.
echo [1/2] Building CLI (qoder-sm.exe) v%VERSION%...
go build -ldflags "%LDFLAGS%" -o build\bin\qoder-sm.exe .
if %errorlevel% equ 0 (echo   OK) else (echo   FAILED)
echo.
echo [2/2] Building Web Server (qoder-web.exe) v%VERSION%...
go build -ldflags "%LDFLAGS%" -o build\bin\qoder-web.exe .\cmd\web\
if %errorlevel% equ 0 (echo   OK) else (echo   FAILED)
echo.
echo Done! Check build\bin\
