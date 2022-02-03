
Set-Variable -Name "PRODUCT_NAME" -Value "Golden Point"
Set-Variable -Name "PRODUCT_VERSION" -Value "1.2.17"

Write-Host "=== Golden Point Compile Script - MS-Windows ==="

# Section 1. Generate assets...
#
Write-Host "Step 1. Generate assets..."
$invokeExpressionOptions = @{
    Command = "go generate"
}
Invoke-Expression @invokeExpressionOptions

# Section 2. Compile executables...
#
Write-Host "Step 2. Compile executables..."
$invokeExpressionOptions = @{
    Command = "go get -v -u"
}
#[Environment]::GetEnvironmentVariables()
Invoke-Expression @invokeExpressionOptions

# Section 2.1. Compile X86_64 executable...
#
Write-Host "Step 2.1. Compile X86_64 executable..."
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

# Section 2.2. Compile X86 executable...
#
Write-Host "Step 2.2. Compile X86 executable..."
$Env:GOOS = "windows"
$Env:GOARCH = "386"
$Env:CGO_ENABLED = "1"
$Env:CC = "gcc.exe"
$Env:CXX = "g++.exe"
$ARCH = "386"
$invokeExpressionOptions = @{
    Command = "go build -o golden-windows-$ARCH.exe"
}
#[Environment]::GetEnvironmentVariables()
Invoke-Expression @invokeExpressionOptions

# Section 3. Make ZIP portable distribution package...
#
Write-Host "Step 3. Make ZIP portable distribution package..."
$TimeStamp = $(Get-Date -Format 'yyyyMMddHHmmtt')
$compressArchiveOptions = @{
    CompressionLevel = "Optimal"
    LiteralPath = "golden-windows-amd64.exe", "golden-windows-386.exe", "ChangeLog", "LICENSE", "README.md", "docs"
    DestinationPath = "Golden-Point-${PRODUCT_VERSION}-${TimeStamp}.zip"
}
Compress-Archive @compressArchiveOptions
