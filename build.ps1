# QoderSessionManager Windows Build Script
param(
    [string]$Target = "all"
)

$Version = "1.0.0"
$BuildTime = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
$GitCommit = git rev-parse --short HEAD 2>$null
if (-not $GitCommit) { $GitCommit = "unknown" }

$LDFlags = "-X 'qoder-sm/pkg/qoder.Version=$Version' -X 'qoder-sm/pkg/qoder.BuildTime=$BuildTime' -X 'qoder-sm/pkg/qoder.GitCommit=$GitCommit'"

$BuildDir = "build\bin"
if (!(Test-Path $BuildDir)) { New-Item -ItemType Directory -Path $BuildDir -Force | Out-Null }

function Build-CLI {
    Write-Host "Building CLI (qoder-sm.exe) v$Version..." -ForegroundColor Cyan
    go build -ldflags $LDFlags -o "$BuildDir\qoder-sm.exe" .
    if ($LASTEXITCODE -eq 0) { Write-Host "  OK" -ForegroundColor Green } else { Write-Host "  FAILED" -ForegroundColor Red }
}

function Build-Web {
    Write-Host "Building Web Server (qoder-web.exe) v$Version..." -ForegroundColor Cyan
    go build -ldflags $LDFlags -o "$BuildDir\qoder-web.exe" .\cmd\web\
    if ($LASTEXITCODE -eq 0) { Write-Host "  OK" -ForegroundColor Green } else { Write-Host "  FAILED" -ForegroundColor Red }
}

switch ($Target) {
    "cli"  { Build-CLI }
    "web"  { Build-Web }
    "all"  { Build-CLI; Build-Web }
}

Write-Host "`nBuild complete! Binaries in: $BuildDir" -ForegroundColor Yellow
