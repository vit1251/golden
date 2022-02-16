
Set-Variable -Name "PRODUCT_NAME" -Value "Golden Point"
Set-Variable -Name "PRODUCT_VERSION" -Value "1.2.17"

Write-Host "=== Golden Point Compile Script - MS-Windows ==="

# Section 1. Setup depenencies...
#
Write-Host "Step 1. Setup depenencies..."
$invokeExpressionOptions = @{
    Command = "go get -v -u"
}
#[Environment]::GetEnvironmentVariables()
Invoke-Expression @invokeExpressionOptions

# Section 2. Generate assets...
#
Write-Host "Step 2. Generate assets..."
$invokeExpressionOptions = @{
    Command = "go generate"
}
Invoke-Expression @invokeExpressionOptions

# Section 3. Compile X86_64 executable...
#
Write-Host "Step 3. Compile X86_64 executable..."
$Env:GOOS = "windows"
$Env:GOARCH = "amd64"
$Env:CGO_ENABLED = "1"
$Env:CC = "x86_64-w64-mingw32-gcc.exe"
$Env:CXX = "x86_64-w64-mingw32-g++.exe"
$ARCH = "amd64"
$invokeExpressionOptions = @{
    Command = "go build -o golden-windows-$ARCH.exe"
}
#[Environment]::GetEnvironmentVariables()
Invoke-Expression @invokeExpressionOptions

# Section 4. Make ZIP portable distribution package...
#
Write-Host "Step 4. Make ZIP portable distribution package..."
$TimeStamp = $(Get-Date -Format 'yyyyMMddHHmmtt')
$compressArchiveOptions = @{
    CompressionLevel = "Optimal"
    LiteralPath = "golden-windows-amd64.exe", "ChangeLog", "LICENSE", "README.md", "docs"
    DestinationPath = "Golden-Point-amd64-${PRODUCT_VERSION}-${TimeStamp}.zip"
}
Compress-Archive @compressArchiveOptions
