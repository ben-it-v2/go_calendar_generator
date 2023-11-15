if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    # Command "go" isn't found
    # Execute powershell script to install Golang
    $installScriptPath = "./install.ps1"
    Invoke-Expression -Command $installScriptPath
}

# Clean go project
go clean

# Build calendar.exe file
go build
