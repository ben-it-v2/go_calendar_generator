if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    # Command "go" isn't found
    # Install Golang
    Write-Host "Installing Golang..."

    # Create the directory
    # Write-Host "Creating Go directory..."
    $goPath = "$env:USERPROFILE\go"
    # New-Item -ItemType Directory -Path $goPath | Out-Null
    # Write-Host "Created Go directory: $goPath"

    # Download ZIP file
    # Go v1.20.2 AMD64
    Write-Host "Downloading Golang ZIP..."
    $url = "https://go.dev/dl/go1.20.2.windows-amd64.zip"
    $output = "$env:USERPROFILE\Downloads\go1.20.2.windows-amd64.zip"
    Invoke-WebRequest -Uri $url -OutFile $output
    Write-Host "Downloaded Golang ZIP"

    # Unzip the archive
    Write-Host "Unzipping Golang ZIP..."
    Expand-Archive -Path $output -DestinationPath $env:USERPROFILE
    Write-Host "Unzipped Golang ZIP: $goPath"

    # Setup Environment Variables
    Write-Host "Setup Environmenbt Varibles..."
    $goBin = "$goPath\bin"
    $goSrc = "$env:USERPROFILE\gosrc"
    # [Environment]::SetEnvironmentVariable("GOROOT", $goPath, "User")
    # [Environment]::SetEnvironmentVariable("GOPATH", $goSrc, "User")
    # [Environment]::SetEnvironmentVariable("Path", "$goBin;$Env:Path", "User")
    $env:GOROOT = $goPath # Or the directory where you extracted Go
    $env:GOPATH = "$goSrc"
    $env:PATH += ";$goBin"
} else {
    # Command "go" is found
    # Golang is already installed
    Write-Host "Golang is already installed."
}

# Display Golang version
go version
