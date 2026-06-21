# QoderSessionManager Windows Build Script
param(
    [string]$Target = "all"
)

$BuildDir = "build\bin"
if (!(Test-Path $BuildDir)) { New-Item -ItemType Directory -Path $BuildDir -Force | Out-Null }

function Build-CLI {
    Write-Host "Building CLI (qoder-sm.exe)..." -ForegroundColor Cyan
    go build -o "$BuildDir\qoder-sm.exe" .
    if ($LASTEXITCODE -eq 0) { Write-Host "  OK" -ForegroundColor Green } else { Write-Host "  FAILED" -ForegroundColor Red }
}

function Build-Web {
    Write-Host "Building Web Server (qoder-web.exe)..." -ForegroundColor Cyan
    go build -o "$BuildDir\qoder-web.exe" .\cmd\web\
    if ($LASTEXITCODE -eq 0) { Write-Host "  OK" -ForegroundColor Green } else { Write-Host "  FAILED" -ForegroundColor Red }
}

switch ($Target) {
    "cli"  { Build-CLI }
    "web"  { Build-Web }
    "all"  { Build-CLI; Build-Web }
}

Write-Host "`nBuild complete! Binaries in: $BuildDir" -ForegroundColor Yellow
