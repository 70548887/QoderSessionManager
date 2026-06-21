@echo off
echo Building QoderSessionManager for Windows...
if not exist build\bin mkdir build\bin
echo.
echo [1/2] Building CLI (qoder-sm.exe)...
go build -o build\bin\qoder-sm.exe .
if %errorlevel% equ 0 (echo   OK) else (echo   FAILED)
echo.
echo [2/2] Building Web Server (qoder-web.exe)...
go build -o build\bin\qoder-web.exe .\cmd\web\
if %errorlevel% equ 0 (echo   OK) else (echo   FAILED)
echo.
echo Done! Check build\bin\
